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
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//
// Performs  a POST into a URL
//
func PostForm(__url string, __body url.Values) ([]byte, error) {
	res, err := http.PostForm(__url, __body)
	if err != nil {
		return nil, err
	} else {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		} else {
			return body, nil
		}
	}
}

//
// Performs a POST in a URL
//
func Post(__url string) ([]byte, error) {

	return PostForm(__url, nil)

}

//
// Performas a POST in a URL with a body.
//
func PostBody(__url string, __body []byte) ([]byte, error) {

	_, _, _resp, _err := Perform(POST, __url, __body, nil)
	return _resp, _err

}

//
// Performs a GET in a URL
//
func Get(__url string) ([]byte, error) {
	res, err := http.Get(__url)
	if err != nil {
		return nil, err
	} else {
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		} else {
			return body, nil
		}
	}
}

//
// Performs a PUT in a URL
//
func Put(__url string) ([]byte, error) {

	_, _, _resp, _err := Perform(PUT, __url, nil, nil)
	return _resp, _err

}

//
// Performs a DELETE in a URL
//
func Delete(__url string) ([]byte, error) {

	_, _, _resp, _err := Perform(DELETE, __url, nil, nil)
	return _resp, _err

}

//
// Generic HTTP perform method.
//
func Perform(__command Method, __url string, __body []byte, __header http.Header) (string, int, []byte, error) {

	var _req *http.Request

	_req, _err := http.NewRequest(string(__command), __url, bytes.NewBuffer(__body))
	if _err != nil {
		return "", -1, nil, _err
	} else {

		// Add headers if there exists
		if __header != nil {
			_req.Header = __header
		}

		_client := &http.Client{Timeout: DEFAULT_TIMEOUT}
		_resp, _err := _client.Do(_req)
		if _err != nil {
			log.Printf("[SC] Http Error: %s", _err.Error())
			return "", -1, nil, _err
		} else {
			defer _resp.Body.Close()
			_status := _resp.Status
			_statusCode := _resp.StatusCode
			_body, _err := ioutil.ReadAll(_resp.Body)
			if _err != nil {
				log.Printf("[SC] Http Error: %s", _err.Error())
				return _status, _statusCode, nil, _err
			} else {
				return _status, _statusCode, _body, nil
			}
		}

	}

}
