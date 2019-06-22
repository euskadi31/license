// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

// License struct
type License struct {
	ID         string  `json:"id"`
	Filename   string  `json:"filename"`
	Confidence float64 `json:"confidence"`
}
