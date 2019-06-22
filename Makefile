BUILD=build

.PHONY: release
release:
	@echo "Release v$(version)"
	@git pull
	@git checkout master
	@git pull
	@git checkout develop
	@git flow release start $(version)
	@git flow release finish $(version) -p -m "Release v$(version)"
	@git checkout develop
	@echo "Release v$(version) finished."

.PHONY: all
all: coverage.out

coverage.out: $(shell find . -type f -print | grep -v vendor | grep "\.go")
	@go test -race -cover -covermode=atomic -coverprofile ./coverage.out.tmp ./...
	@cat ./coverage.out.tmp | grep -v '.pb.go' | grep -v 'mock_' > ./coverage.out
	@rm ./coverage.out.tmp

.PHONY: test
test: coverage.out

.PHONY: cover
cover: coverage.out
	@echo ""
	@go tool cover -func ./coverage.out

.PHONY: cover-html
cover-html: coverage.out
	@go tool cover -html=./coverage.out

.PHONY: benchmark
benchmark:
	@go test -bench=. ./...

.PHONY: clean
clean:
	@rm ./coverage.out
	@go clean -i ./...

.PHONY: generate
generate:
	@go generate ./...

.PHONY: lint
lint:
	@golangci-lint run ./...

.PHONY: docs
docs:
	@mkdir -p /tmp/tmpgoroot/doc
	@rm -rf /tmp/tmpgopath/src/github.com/euskadi31/go-eventemitter
	@mkdir -p /tmp/tmpgopath/src/github.com/euskadi31/go-eventemitter
	@tar -c --exclude='.git' --exclude='tmp' . | tar -x -C /tmp/tmpgopath/src/github.com/euskadi31/go-eventemitter
	@echo -e "open http://localhost:6060/pkg/github.com/euskadi31/go-eventemitter\n"
	@GOROOT=/tmp/tmpgoroot/ GOPATH=/tmp/tmpgopath/ godoc -http=localhost:6060

pkg/license/template/bindata.go: $(shell find licenses -type f -name *.tpl)
	@echo "Bin data..."
	@go-bindata -pkg template -o $@ licenses/

${BUILD}/license: $(shell find . -type f -print | grep -v vendor | grep "\.go") pkg/license/template/bindata.go
	@echo "Building license..."
	@go generate ./cmd/license/
	@go build -o $@ ./cmd/license/

.PHONY: build
build: ${BUILD}/license

run-license: ${BUILD}/license
	@echo "Running license..."
	@./$<

.PHONY: run
run: run-license
