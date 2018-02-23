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

//
// Based on:
// http://www.iana.org/assignments/http-status-codes
//
type StatusCode struct {
	_Code int
	_Desc string
	_RFC  string
}

func (__obj *StatusCode) Code() int {
	return __obj._Code
}

func (__obj *StatusCode) Desc() string {
	return __obj._Desc
}

func (__obj *StatusCode) RFC() string {
	return __obj._RFC
}

type StatusCodeRange struct {
	_Min  int
	_Max  int
	_Desc string
	_RFC  string
}

func (__obj *StatusCodeRange) Min() int {
	return __obj._Min
}

func (__obj *StatusCodeRange) Max() int {
	return __obj._Max
}

func (__obj *StatusCodeRange) Desc() string {
	return __obj._Desc
}

func (__obj *StatusCodeRange) RFC() string {
	return __obj._RFC
}

var HTTP_1xx = StatusCodeRange{_Min: 100, _Max: 199, _Desc: "1xx"}
var HTTP_100 = StatusCode{_Code: 100, _Desc: "Continue", _RFC: "[RFC7231, Section 6.2.1]"}
var HTTP_101 = StatusCode{_Code: 101, _Desc: "Switching Protocols", _RFC: "[RFC7231, Section 6.2.2]"}
var HTTP_102 = StatusCode{_Code: 102, _Desc: "Processing", _RFC: "[RFC2518]"}
var HTTP_103 = StatusCode{_Code: 103, _Desc: "Early Hints", _RFC: "[RFC8297]"}

var HTTP_2xx = StatusCodeRange{_Min: 200, _Max: 299, _Desc: "2xx"}
var HTTP_200 = StatusCode{_Code: 200, _Desc: "OK", _RFC: "[RFC7231, Section 6.3.1]"}
var HTTP_201 = StatusCode{_Code: 201, _Desc: "Created", _RFC: "[RFC7231, Section 6.3.2]"}
var HTTP_202 = StatusCode{_Code: 202, _Desc: "Accepted", _RFC: "[RFC7231, Section 6.3.3]"}
var HTTP_203 = StatusCode{_Code: 203, _Desc: "Non-Authoritative Information", _RFC: "[RFC7231, Section 6.3.4]"}
var HTTP_204 = StatusCode{_Code: 204, _Desc: "No Content", _RFC: "[RFC7231, Section 6.3.5]"}
var HTTP_205 = StatusCode{_Code: 205, _Desc: "Reset Content", _RFC: "[RFC7231, Section 6.3.6]"}
var HTTP_206 = StatusCode{_Code: 206, _Desc: "Partial Content", _RFC: "[RFC7233, Section 4.1]"}
var HTTP_207 = StatusCode{_Code: 207, _Desc: "Multi-Status", _RFC: "[RFC4918]"}
var HTTP_208 = StatusCode{_Code: 208, _Desc: "Already Reported", _RFC: "[RFC5842]"}
var HTTP_226 = StatusCode{_Code: 226, _Desc: "IM Used", _RFC: "[RFC3229]"}

var HTTP_3xx = StatusCodeRange{_Min: 300, _Max: 399, _Desc: "3xx"}
var HTTP_300 = StatusCode{_Code: 300, _Desc: "Multiple Choices", _RFC: "[RFC7231, Section 6.4.1]"}
var HTTP_301 = StatusCode{_Code: 301, _Desc: "Moved Permanently", _RFC: "[RFC7231, Section 6.4.2]"}
var HTTP_302 = StatusCode{_Code: 302, _Desc: "Found", _RFC: "[RFC7231, Section 6.4.3]"}
var HTTP_303 = StatusCode{_Code: 303, _Desc: "See Other", _RFC: "[RFC7231, Section 6.4.4]"}
var HTTP_304 = StatusCode{_Code: 304, _Desc: "Not Modified", _RFC: "[RFC7232, Section 4.1]"}
var HTTP_305 = StatusCode{_Code: 305, _Desc: "Use Proxy", _RFC: "[RFC7231, Section 6.4.5]"}
var HTTP_306 = StatusCode{_Code: 306, _Desc: "(Unused)", _RFC: "[RFC7231, Section 6.4.6]"}
var HTTP_307 = StatusCode{_Code: 307, _Desc: "Temporary Redirect", _RFC: "[RFC7231, Section 6.4.7]"}
var HTTP_308 = StatusCode{_Code: 308, _Desc: "Permanent Redirect", _RFC: "[RFC7538]"}

var HTTP_4xx = StatusCodeRange{_Min: 400, _Max: 499, _Desc: "4xx"}
var HTTP_400 = StatusCode{_Code: 400, _Desc: "Bad Request", _RFC: "[RFC7231, Section 6.5.1]"}
var HTTP_401 = StatusCode{_Code: 401, _Desc: "Unauthorized", _RFC: "[RFC7235, Section 3.1]"}
var HTTP_402 = StatusCode{_Code: 402, _Desc: "Payment Required", _RFC: "[RFC7231, Section 6.5.2]"}
var HTTP_403 = StatusCode{_Code: 403, _Desc: "Forbidden", _RFC: "[RFC7231, Section 6.5.3]"}
var HTTP_404 = StatusCode{_Code: 404, _Desc: "Not Found", _RFC: "[RFC7231, Section 6.5.4]"}
var HTTP_405 = StatusCode{_Code: 405, _Desc: "Method Not Allowed", _RFC: "[RFC7231, Section 6.5.5]"}
var HTTP_406 = StatusCode{_Code: 406, _Desc: "Not Acceptable", _RFC: "[RFC7231, Section 6.5.6]"}
var HTTP_407 = StatusCode{_Code: 407, _Desc: "Proxy Authentication Required", _RFC: "[RFC7235, Section 3.2]"}
var HTTP_408 = StatusCode{_Code: 408, _Desc: "Request Timeout", _RFC: "[RFC7231, Section 6.5.7]"}
var HTTP_409 = StatusCode{_Code: 409, _Desc: "Conflict", _RFC: "[RFC7231, Section 6.5.8]"}
var HTTP_410 = StatusCode{_Code: 410, _Desc: "Gone", _RFC: "[RFC7231, Section 6.5.9]"}
var HTTP_411 = StatusCode{_Code: 411, _Desc: "Length Required", _RFC: "[RFC7231, Section 6.5.10]"}
var HTTP_412 = StatusCode{_Code: 412, _Desc: "Precondition Failed", _RFC: "[RFC7232, Section 4.2][RFC8144, Section 3.2]"}
var HTTP_413 = StatusCode{_Code: 413, _Desc: "Payload Too Large", _RFC: "[RFC7231, Section 6.5.11]"}
var HTTP_414 = StatusCode{_Code: 414, _Desc: "URI Too Long", _RFC: "[RFC7231, Section 6.5.12]"}
var HTTP_415 = StatusCode{_Code: 415, _Desc: "Unsupported Media Type", _RFC: "[RFC7231, Section 6.5.13][RFC7694, Section 3]"}
var HTTP_416 = StatusCode{_Code: 416, _Desc: "Range Not Satisfiable", _RFC: "[RFC7233, Section 4.4]"}
var HTTP_417 = StatusCode{_Code: 417, _Desc: "Expectation Failed", _RFC: "[RFC7231, Section 6.5.14]"}
var HTTP_421 = StatusCode{_Code: 421, _Desc: "Misdirected Request", _RFC: "[RFC7540, Section 9.1.2]"}
var HTTP_422 = StatusCode{_Code: 422, _Desc: "Unprocessable Entity", _RFC: "[RFC4918]"}
var HTTP_423 = StatusCode{_Code: 423, _Desc: "Locked", _RFC: "[RFC4918]"}
var HTTP_424 = StatusCode{_Code: 424, _Desc: "Failed Dependency", _RFC: "[RFC4918]"}
var HTTP_426 = StatusCode{_Code: 426, _Desc: "Upgrade Required", _RFC: "[RFC7231, Section 6.5.15]"}
var HTTP_428 = StatusCode{_Code: 428, _Desc: "Precondition Required", _RFC: "[RFC6585]"}
var HTTP_429 = StatusCode{_Code: 429, _Desc: "Too Many Requests", _RFC: "[RFC6585]"}
var HTTP_431 = StatusCode{_Code: 431, _Desc: "Request Header Fields Too Large", _RFC: "[RFC6585]"}
var HTTP_451 = StatusCode{_Code: 451, _Desc: "Unavailable For Legal Reasons", _RFC: "[RFC7725]"}

var HTTP_5xx = StatusCodeRange{_Min: 500, _Max: 599, _Desc: "5xx"}
var HTTP_500 = StatusCode{_Code: 500, _Desc: "Internal Server Error", _RFC: "[RFC7231, Section 6.6.1]"}
var HTTP_501 = StatusCode{_Code: 501, _Desc: "Not Implemented", _RFC: "[RFC7231, Section 6.6.2]"}
var HTTP_502 = StatusCode{_Code: 502, _Desc: "Bad Gateway", _RFC: "[RFC7231, Section 6.6.3]"}
var HTTP_503 = StatusCode{_Code: 503, _Desc: "Service Unavailable", _RFC: "[RFC7231, Section 6.6.4]"}
var HTTP_504 = StatusCode{_Code: 504, _Desc: "Gateway Timeout", _RFC: "[RFC7231, Section 6.6.5]"}
var HTTP_505 = StatusCode{_Code: 505, _Desc: "HTTP Version Not Supported", _RFC: "[RFC7231, Section 6.6.6]"}
var HTTP_506 = StatusCode{_Code: 506, _Desc: "Variant Also Negotiates", _RFC: "[RFC2295]"}
var HTTP_507 = StatusCode{_Code: 507, _Desc: "Insufficient Storage", _RFC: "[RFC4918]"}
var HTTP_508 = StatusCode{_Code: 508, _Desc: "Loop Detected", _RFC: "[RFC5842]"}
var HTTP_510 = StatusCode{_Code: 510, _Desc: "Not Extended", _RFC: "[RFC2774]"}
var HTTP_511 = StatusCode{_Code: 511, _Desc: "Network Authentication Required", _RFC: "[RFC6585]"}
