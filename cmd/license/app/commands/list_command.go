// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"

	"github.com/euskadi31/license/pkg/license/identifier"
	"github.com/urfave/cli"
)

// ListCommand constructor
func ListCommand() cli.Command {
	return cli.Command{
		Name:   "list",
		Usage:  "Shows a list of available licenses",
		Action: listAction,
	}
}

func listAction(c *cli.Context) error {
	for _, id := range identifier.List() {
		fmt.Fprintf(c.App.Writer, " * %s\n", id)
	}

	return nil
}
