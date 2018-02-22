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
	"encoding/json"
)

//
// Json_Unmarshal_Object utilitary method.
//
func Json_Unmarshal_Object(data []byte) (map[string]*json.RawMessage, error) {

	var objmap map[string]*json.RawMessage
	e := json.Unmarshal(data, &objmap)
	return objmap, e

}

//
// Json_Unmarshal_Array utilitary method.
//
func Json_Unmarshal_Array(data []byte) ([]map[string]*json.RawMessage, error) {

	var arrmap []map[string]*json.RawMessage
	e := json.Unmarshal(data, &arrmap)
	return arrmap, e

}
