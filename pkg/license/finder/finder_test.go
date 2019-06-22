// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLicenseFileScore(t *testing.T) {
	assertions := []struct {
		value    string
		expected float64
	}{
		{
			value:    "LICENSE",
			expected: 1.0,
		},
		{
			value:    "LICENSE.md",
			expected: 0.95,
		},
		{
			value:    "COPYING",
			expected: 0.90,
		},
		{
			value:    "COPYING.html",
			expected: 0.85,
		},
		{
			value:    "LICENSE-MIT",
			expected: 0.70,
		},
		{
			value:    "COPYING-MIT",
			expected: 0.65,
		},
		{
			value:    "MIT-LICENSE-MIT",
			expected: 0.60,
		},
		{
			value:    "MIT-COPYING",
			expected: 0.55,
		},
		{
			value:    "OFL.md",
			expected: 0.50,
		},
		{
			value:    "OFL",
			expected: 0.40,
		},
		{
			value:    "PATENTS",
			expected: 0.35,
		},
		{
			value:    "PATENTS.txt",
			expected: 0.30,
		},
		{
			value:    "foo",
			expected: 0.0,
		},
	}

	for _, assertion := range assertions {
		assert.Equal(t, assertion.expected, LicenseFileScore(assertion.value))
	}
}

func TestIsLicenseFile(t *testing.T) {
	assertions := []struct {
		value    string
		expected bool
	}{
		{
			value:    "LICENSE",
			expected: true,
		},
		{
			value:    "LICENSE.md",
			expected: true,
		},
		{
			value:    "COPYING",
			expected: true,
		},
		{
			value:    "COPYING.html",
			expected: true,
		},
		{
			value:    "LICENSE-MIT",
			expected: true,
		},
		{
			value:    "COPYING-MIT",
			expected: true,
		},
		{
			value:    "MIT-LICENSE-MIT",
			expected: true,
		},
		{
			value:    "MIT-COPYING",
			expected: true,
		},
		{
			value:    "OFL.md",
			expected: true,
		},
		{
			value:    "OFL",
			expected: true,
		},
		{
			value:    "PATENTS",
			expected: true,
		},
		{
			value:    "PATENTS.txt",
			expected: true,
		},
		{
			value:    "foo",
			expected: false,
		},
	}

	for _, assertion := range assertions {
		assert.Equal(t, assertion.expected, IsLicenseFile(assertion.value))
	}
}
