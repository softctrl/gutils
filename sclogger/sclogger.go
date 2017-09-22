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
package sclogger

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	_TAG_         = "*"
	_LOG_HTTP_FMT = "%s\t%s\t%s\t%s\t%s"
)

const (
	_ERROR_FMT   = "[E][%s] - %s"
	_INFO_FMT    = "[I][%s] - %s"
	_WARNING_FMT = "[W][%s] - %s"
)

//
// Handle logging into a http server.
//
func HttpHandlerLogger(inner http.Handler, name string) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		inner.ServeHTTP(w, r)

		I(_TAG_, _LOG_HTTP_FMT,
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			name,
			time.Since(start))

	})

}

//
// Prints an error message into the current log.
//
func E(__tag, __fmt string, v ...interface{}) {

	log.Printf(_ERROR_FMT, __tag, fmt.Sprintf(__fmt, v...))

}

//
// Prints a warning message into the current log.
//
func W(__tag, __fmt string, v ...interface{}) {

	log.Printf(_WARNING_FMT, __tag, fmt.Sprintf(__fmt, v...))

}

//
// Prints an info message into the current log.
//
func I(__tag, __fmt string, v ...interface{}) {

	log.Printf(_INFO_FMT, __tag, fmt.Sprintf(__fmt, v...))

}

//
//
//
func Fatal(v ...interface{}) {

	log.Fatal(v...)

}

//
//
//
func Printf(format string, v ...interface{}) {

	log.Printf(format, v...)

}
