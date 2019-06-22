// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/euskadi31/license/cmd/license/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		panic(err)
	}

	if err := application.Run(os.Args); err != nil {
		panic(err)
	}
}
