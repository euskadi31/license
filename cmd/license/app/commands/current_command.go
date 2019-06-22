// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/euskadi31/license/pkg/license/extractor"
	"github.com/euskadi31/license/pkg/license/finder"
	"github.com/google/licenseclassifier"
	"github.com/urfave/cli"
)

// CurrentCommand constructor
func CurrentCommand() cli.Command {
	return cli.Command{
		Name:   "current",
		Usage:  "Shows current license",
		Action: currentAction,
	}
}

func findLicenseFile(files []os.FileInfo) string {
	for _, info := range files {
		if info.IsDir() {
			continue
		}

		if finder.IsLicenseFile(info.Name()) {
			return info.Name()
		}
	}

	return ""
}

func currentAction(c *cli.Context) error {
	threshold := licenseclassifier.DefaultConfidenceThreshold

	root := c.Args().First()

	if root == "" {
		root = "."
	}

	root = strings.TrimRight(root, "/")

	files, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	filename := findLicenseFile(files)

	if filename == "" {
		return errors.New("License not found")
	}

	filename = root + "/" + filename

	classifier, err := extractor.New(threshold)
	if err != nil {
		return err
	}

	errs := classifier.Extracts([]string{filename}, false)
	if len(errs) > 0 {
		for _, err := range errs {
			return err
		}
	}

	licenses := classifier.GetResults()

	if len(licenses) == 0 {
		return errors.New("Couldn't classify license(s)")
	}

	sort.Sort(licenses)

	for _, license := range licenses {
		fmt.Fprintf(c.App.Writer, "License: %s\n", license.Name)

		break
	}

	return nil
}
