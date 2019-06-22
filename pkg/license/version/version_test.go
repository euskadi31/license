// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package version

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	info := Get()

	assert.Equal(t, runtime.Version(), info.GoVersion)
}

func TestParse(t *testing.T) {
	v, err := Parse("v1.0.1")
	assert.NoError(t, err)

	assert.Equal(t, "1.0.1", v.String())
}

func TestParseWithError(t *testing.T) {
	v, err := Parse("v1.")
	assert.Error(t, err)
	assert.Nil(t, v)
}
