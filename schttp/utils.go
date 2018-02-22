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
package schttp

import (
	"net/http"
	str "strings"
	"time"
)

type Method string

type Methods []Method

// Common HTTP methods.
//
// Unless otherwise noted, these are defined in RFC 7231 section 4.3.
const (
	GET     = Method(http.MethodGet)
	HEAD    = Method(http.MethodHead)
	POST    = Method(http.MethodPost)
	PUT     = Method(http.MethodPut)
	PATCH   = Method(http.MethodPatch)
	DELETE  = Method(http.MethodDelete)
	CONNECT = Method(http.MethodConnect)
	OPTIONS = Method(http.MethodOptions)
	TRACE   = Method(http.MethodTrace)
)

var ALL_METHODS = Methods{GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE}

//
// MethodFromString returns a Method instance from an informed string.
//
func MethodFromString(__method string) Method {
	return Method(str.ToUpper(__method))
}

const DEFAULT_TIMEOUT = time.Second * 200
