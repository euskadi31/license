// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/euskadi31/license/pkg/license/extractor"
	"github.com/euskadi31/license/pkg/license/finder"
	"github.com/euskadi31/license/pkg/license/types"
	"github.com/google/licenseclassifier"
	"github.com/urfave/cli"
)

// ExtractCommand constructor
func ExtractCommand() cli.Command {
	return cli.Command{
		Name:  "extract",
		Usage: "Extract all license in your project and vendor libraries",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "format, f",
				Value: "json",
				Usage: "license format (json, html, text)",
			},
			cli.StringFlag{
				Name:  "output, o",
				Value: ".",
				Usage: "path to output file",
			},
			/*cli.StringFlag{
				Name:  "html-template",
				Value: "default",
				Usage: "the template name",
			},*/
		},
		Action: extractAction,
	}
}

func extractAction(c *cli.Context) error {
	threshold := licenseclassifier.DefaultConfidenceThreshold

	format := c.String("format")
	output := c.String("output")

	root := c.Args().First()

	if root == "" {
		root = "."
	}

	if output == "" {
		output = "."
	}

	output = strings.TrimRight(output, "/")

	filenames := []string{}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if finder.IsLicenseFile(path) {
			filenames = append(filenames, path)
		}

		return nil
	})
	if err != nil {
		return err
	}

	classifier, err := extractor.New(threshold)
	if err != nil {
		return err
	}

	errs := classifier.Extracts(filenames, false)
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

	response := []*types.License{}

	stats := map[string]int{}
	for _, r := range licenses {
		if _, ok := stats[r.Name]; !ok {
			stats[r.Name] = 0
		}

		stats[r.Name]++

		response = append(response, &types.License{
			ID:         r.Name,
			Filename:   r.Filename,
			Confidence: r.Confidence,
		})
	}

	fmt.Fprintf(c.App.Writer, "License found:\n")

	p := make(PairList, len(stats))

	i := 0
	for k, v := range stats {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	for _, key := range p {
		fmt.Fprintf(c.App.Writer, " * %s: %d\n", key.Key, key.Value)
	}

	switch format {
	case "json":
		content, err := json.MarshalIndent(response, "", " ")
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(output+"/licenses.json", content, 0644); err != nil {
			return err
		}
	}

	return nil
}

// Pair struct
type Pair struct {
	Key   string
	Value int
}

// PairList A slice of pairs that implements sort.Interface to sort by values
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }
