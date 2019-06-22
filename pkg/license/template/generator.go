// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package template

import (
	"bytes"
	"html/template"
)

// Generate license from identifier and Context
func Generate(identifier string, data *Context) ([]byte, error) {
	content, err := Asset("licenses/" + identifier + ".tpl")
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("identifier").Parse(string(content))
	if err != nil {
		return nil, err
	}
	var tpl bytes.Buffer

	if err := tmpl.Execute(&tpl, data); err != nil {
		return nil, err
	}

	return tpl.Bytes(), nil
}
