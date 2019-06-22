// Copyright 2019 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package types

// Project struct
type Project struct {
	Name    string   `json:"name"`
	License *License `json:"license"`
}
