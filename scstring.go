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
	b "bytes"
	f "fmt"
	str "strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	JOINED_QUOTED_VALUES_FMT = "[%s]"
	EMPTY_ARRAY              = "[]"
	QUOTE                    = '"'
	COMMA                    = ','
)

//
//
//
func JoinQuoted(__values []string, sep string) string {

	_len := len(__values)
	if _len > 0 {
		var _buff b.Buffer
		_buff.WriteString(QuotedString(__values[0]))
		for _idx := 1; _idx < _len; _idx++ {
			_buff.WriteByte(COMMA)
			_buff.WriteString(QuotedString(__values[_idx]))
		}
		return f.Sprintf(JOINED_QUOTED_VALUES_FMT, _buff.String())
	} else {
		return EMPTY_ARRAY
	}

}

//
//
//
func Quoted(__value []byte) []byte {

	var _buff b.Buffer
	_buff.WriteByte(QUOTE)
	_buff.Write(__value)
	_buff.WriteByte(QUOTE)
	return _buff.Bytes()

}

//
//
//
func QuotedString(__value string) string {

	var _buff b.Buffer
	_buff.WriteByte(QUOTE)
	_buff.WriteString(__value)
	_buff.WriteByte(QUOTE)
	return _buff.String()

}

//
// Based on: https://gist.github.com/tkrajina/d423e9b9fc2e72d63072
//
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

//
// Remove all accents from an informed byte array
// Based on: https://gist.github.com/tkrajina/d423e9b9fc2e72d63072
//
func RemoveAccents(__value []byte) []byte {

	__result := make([]byte, len(__value))
	_transform := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	_, _, _err := _transform.Transform(__result, __value, true)
	if _err != nil {
		panic(_err)
	}
	return __result

}

//
// Remove all accents from an informed string
//
func RemoveAccentsString(__value string) string {
	return string(RemoveAccents([]byte(__value)))
}

//
// Remove all duplicated spaces from an informed string
//
//
func RemoveDuplicatedSpaces(__value string) string {
	return DUPLICATED_SPACES_PATTERN.ReplaceAllString(__value, " ")
}

//
//
//
func RemoveNonWords(__value string) string {
	return RemoveDuplicatedSpaces(NON_WORDS_PATTERN.ReplaceAllString(__value, " "))
}

//
// Validate when an informed string is empty.
//
func IsEmpty(__value string) bool {
	return (len(str.TrimSpace(__value)) > 0)
}
