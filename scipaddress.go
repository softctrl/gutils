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

package scgotils

import (
	"strconv"
	"strings"
)

const (
	_256_0 = 1
	_256_1 = 256
	_256_2 = 65536
	_256_3 = 16777216
)

//
//
//
func IsIpAddress(__ip string) bool {
	return IP_ADDRESS_PATTERN.MatchString(__ip)
}

func IntToIp(__ip int) string {
	return ""
}

//
// Based on:
// http://www.0x09.com.br/2013/10/numerical-representation-of-ipv4-address.html
//
func IpToInt(__ip string) int64 {

	if IsIpAddress(__ip) {

		octects := strings.Split(__ip, ".")
		var i_ip int64 = 0

		_i, _ := strconv.Atoi(octects[0])
		i_ip += (int64(_i) * _256_3)

		_i, _ = strconv.Atoi(octects[1])
		i_ip += (int64(_i) * _256_2)

		_i, _ = strconv.Atoi(octects[2])
		i_ip += (int64(_i) * _256_1)

		_i, _ = strconv.Atoi(octects[3])
		i_ip += (int64(_i) * _256_0)

		return i_ip

	} else {
		return -1
	}
}
