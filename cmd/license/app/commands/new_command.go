// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/euskadi31/license/pkg/license/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/euskadi31/license/pkg/license/identifier"
	"github.com/urfave/cli"
)

// NewCommand constructor
func NewCommand() cli.Command {
	return cli.Command{
		Name:  "new",
		Usage: "Generate a license file",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "format, f",
				Value: "",
				Usage: "license format (md)",
			},
		},
		Action: newAction,
	}
}

func newAction(c *cli.Context) error {
	ext := c.String("format")

	context := &template.Context{
		Year: time.Now().Format("2006"),
	}

	id := c.Args().First()

	if id == "" {
		prompt := &survey.Select{
			Message: "Choose a license:",
			Options: identifier.List(),
			Default: "",
		}

		if err := survey.AskOne(prompt, &id); err != nil {
			return err
		}
	}

	if context.CopyrightHolders == "" {
		prompt := &survey.Input{
			Message: "Choose a copyright holders:",
			Help:    "ACME Inc.",
		}

		if err := survey.AskOne(prompt, &context.CopyrightHolders); err != nil {
			return err
		}
	}

	content, err := template.Generate(id, context)
	if err != nil {
		return err
	}

	if ext != "" {
		ext = "." + ext
	}

	if err := ioutil.WriteFile("LICENSE"+ext, content, 0644); err != nil {
		return err
	}

	fmt.Fprintf(c.App.Writer, "LICENSE file created with %s license\n", id)

	return nil
}
