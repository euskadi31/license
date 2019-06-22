// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package app

import (
	"time"

	"github.com/euskadi31/license/cmd/license/app/commands"
	"github.com/euskadi31/license/pkg/license/version"
	"github.com/urfave/cli"
)

// New application
func New() (*cli.App, error) {
	info := version.Get()
	compiled, err := time.Parse("", info.BuildDate)
	if err != nil {
		return nil, err
	}

	year := compiled.Format("2006")

	app := cli.NewApp()
	app.Name = "license"
	app.Version = info.Version
	app.Compiled = compiled
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Axel Etcheverry",
			Email: "axel@etcheverry.biz",
		},
	}
	app.Copyright = "(c) " + year + " Axel Etcheverry"
	app.HelpName = "license tools"
	app.Usage = "License is a tool for creating, extracting and recording all of the license files in your project."
	app.Commands = []cli.Command{
		commands.NewCommand(),
		commands.ListCommand(),
		commands.ExtractCommand(),
		commands.CurrentCommand(),
	}
	app.EnableBashCompletion = true
	app.HideHelp = false
	app.HideVersion = false
	/*app.BashComplete = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "lipstick\nkiss\nme\nlipstick\nringo\n")
	}*/
	app.Action = func(c *cli.Context) error {
		// cli.DefaultAppComplete(c)
		// cli.HandleExitCoder(errors.New("not an exit coder, though"))
		if err := cli.ShowAppHelp(c); err != nil {
			return err
		}
		// cli.ShowCommandCompletions(c, "nope")
		// cli.ShowCommandHelp(c, "also-nope")
		// cli.ShowCompletions(c)
		// cli.ShowSubcommandHelp(c)
		//cli.ShowVersion(c)

		c.App.Setup()

		return nil
	}

	return app, nil
}
