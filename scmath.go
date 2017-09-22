//
// Author:
//  Carlos Timoshenko
//  carlostimoshenkorodrigueslopes@gmail.com
//
//  https://github.com/softctrl
//
// This project is free software; you can redistribute it and/or
// modify it under the terms of the GNU Lesser General Public
// License as published by the Free Software Foundation; either
// version 3 of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// Lesser General Public License for more details.
//
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package gutils

//
// Dim returns the maximum of x-y or 0.
//
func Dim(x, y int) int {
	return Max(x-y, 0)
}

//
// Max returns the larger of x or y.
//
func Max(_x, _y int) int {

	if _x > _y {
		return _x
	} else {
		return _y
	}

}

//
// Min returns the smaller of x or y.
//
func Min(_x, _y int) int {

	if _x < _y {
		return _x
	} else {
		return _y
	}

}
