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

package scgotils

import (
	rgx "regexp"
)

//
// This pattern will only matche a valid integer value.
//
var INTEGER_PATTERN = rgx.MustCompile("^\\d+$")

//
// This pattern will matches duplicated spaces occurencies.
//
var DUPLICATED_SPACES_PATTERN = rgx.MustCompile("\\s+")

//
// It will mathces only symbols there is not from a to Z, 0 to 9, ., space and (").
//
var NON_WORDS_PATTERN = rgx.MustCompile("[^a-zA-Z0-9-. \"]+")

//
// This pattern will matches only valis IPV6 Address.
// Based on:
// https://stackoverflow.com/questions/15875013/extract-ip-addresses-from-strings-using-regex
//
var IP_ADDRESS_PATTERN = rgx.MustCompile("^(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)(\\.(25[0-5]|2[0-4]\\d|[0-1]?\\d?\\d)){3}$")
