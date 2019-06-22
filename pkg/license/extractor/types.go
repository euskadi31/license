// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package extractor

// LicenseType is the assumed type of the unknown license.
type LicenseType struct {
	Filename   string
	Name       string
	Confidence float64
	Offset     int
	Extent     int
}

// LicenseTypes is a list of LicenseType objects.
type LicenseTypes []*LicenseType

func (lt LicenseTypes) Len() int {
	return len(lt)
}

func (lt LicenseTypes) Swap(i, j int) {
	lt[i], lt[j] = lt[j], lt[i]
}

func (lt LicenseTypes) Less(i, j int) bool {
	if lt[i].Confidence > lt[j].Confidence {
		return true
	}

	if lt[i].Confidence < lt[j].Confidence {
		return false
	}

	return lt[i].Filename < lt[j].Filename
}
