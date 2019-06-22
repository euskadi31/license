// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package extractor

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/google/licenseclassifier"
)

// Extractor struct
type Extractor struct {
	results    LicenseTypes
	mu         sync.Mutex
	classifier *licenseclassifier.License
}

// New Extractor
func New(threshold float64) (*Extractor, error) {
	lc, err := licenseclassifier.New(threshold)
	if err != nil {
		return nil, err
	}

	return &Extractor{
		classifier: lc,
	}, nil
}

// ExtractsWithContext runs the license classifier over the given file; ensure that it will respect the timeout in the provided context.
func (e *Extractor) ExtractsWithContext(ctx context.Context, filenames []string, headers bool) (errors []error) {
	done := make(chan bool)
	go func() {
		errors = e.Extracts(filenames, headers)
		done <- true
	}()

	select {
	case <-ctx.Done():
		err := ctx.Err()
		errors = append(errors, err)

		return errors
	case <-done:
		return errors
	}
}

// Extracts licenses
func (e *Extractor) Extracts(filenames []string, headers bool) (errors []error) {
	// Create a pool from which tasks can later be started. We use a pool because the OS limits
	// the number of files that can be open at any one time.
	const numTasks = 1000
	task := make(chan bool, numTasks)
	for i := 0; i < numTasks; i++ {
		task <- true
	}

	errs := make(chan error, len(filenames))

	var wg sync.WaitGroup
	analyze := func(filename string) {
		defer func() {
			wg.Done()
			task <- true
		}()

		if err := e.classifyLicense(filename, headers); err != nil {
			errs <- err
		}
	}

	for _, filename := range filenames {
		wg.Add(1)
		<-task

		go analyze(filename)
	}
	/*
		go func() {
			wg.Wait()
			close(task)
			close(errs)
		}()
	*/

	wg.Wait()
	close(task)
	close(errs)

	for err := range errs {
		errors = append(errors, err)
	}

	return errors
}

// classifyLicense is called by a Go-function to perform the actual
// classification of a license.
func (e *Extractor) classifyLicense(filename string, headers bool) error {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read %q: %v", filename, err)
	}

	matchLoop := func(contents string) {
		for _, m := range e.classifier.MultipleMatch(contents, headers) {
			e.mu.Lock()
			e.results = append(e.results, &LicenseType{
				Filename:   filename,
				Name:       m.Name,
				Confidence: m.Confidence,
				Offset:     m.Offset,
				Extent:     m.Extent,
			})
			e.mu.Unlock()
		}
	}

	//log.Printf("Classifying license(s): %s", filename)
	matchLoop(string(contents))
	/*
		//start := time.Now()
		if lang := language.ClassifyLanguage(filename); lang == language.Unknown {
			matchLoop(string(contents))
		} else {
			//log.Printf("detected language: %v", lang)
			comments := commentparser.Parse(contents, lang)
			for ch := range comments.ChunkIterator() {
				//log.Printf("%q", ch.String())
				matchLoop(ch.String())
			}
		}
	*/
	//log.Printf("Finished Classifying License %q: %v", filename, time.Since(start))

	return nil
}

// GetResults returns the results of the classifications.
func (e *Extractor) GetResults() LicenseTypes {
	return e.results
}
