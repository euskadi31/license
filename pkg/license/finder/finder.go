// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finder

import (
	"path/filepath"
	"regexp"
)

const (
	extensionPattern      = "\\.(md|markdown|txt|html)"
	otherExtensionPattern = ""
	licensePattern        = "(un)?licen[sc]e"
	copyingPattern        = "copy(ing|right)"
	oflPattern            = "ofl"
	patentsPattern        = "patents"
)

// see: github.com/licensee/licensee/lib/licensee/project_files/license_file.rb
var (
	filenameRegexs = map[*regexp.Regexp]float64{
		regexp.MustCompile("(?i)\\A" + licensePattern + "\\z"):                    1,    // LICENSE
		regexp.MustCompile("(?i)\\A" + licensePattern + extensionPattern + "\\z"): 0.95, // LICENSE.md
		regexp.MustCompile("(?i)\\A" + copyingPattern + "\\z"):                    0.90, // COPYING
		regexp.MustCompile("(?i)\\A" + copyingPattern + extensionPattern + "\\z"): 0.85, // COPYING.md
		regexp.MustCompile("(?i)\\A" + licensePattern + "[-_]"):                   0.70, // LICENSE-MIT
		regexp.MustCompile("(?i)\\A" + copyingPattern + "[-_]"):                   0.65, // COPYING-MIT
		regexp.MustCompile("(?i)[-_]" + licensePattern):                           0.60, // MIT-LICENSE-MIT
		regexp.MustCompile("(?i)[-_]" + copyingPattern):                           0.55, // MIT-COPYING
		regexp.MustCompile("(?i)\\A" + oflPattern + extensionPattern + "\\z"):     0.50, // OFL.md
		regexp.MustCompile("(?i)\\A" + oflPattern + "\\z"):                        0.40, // OFL
		regexp.MustCompile("(?i)\\A" + patentsPattern + "\\z"):                    0.35, // PATENTS
		regexp.MustCompile("(?i)\\A" + patentsPattern + extensionPattern + "\\z"): 0.30, // PATENTS.txt
	}
)

// LicenseFileScore return score of file is license
func LicenseFileScore(file string) float64 {
	file = filepath.Base(file)

	for reg, score := range filenameRegexs {
		if reg.MatchString(file) {
			return score
		}
	}

	return 0.0
}

// IsLicenseFile return true if file is license
func IsLicenseFile(file string) bool {
	if LicenseFileScore(file) >= 0.30 {
		return true
	}

	return false
}
