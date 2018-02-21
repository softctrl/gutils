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

//
// Calculate the similarity of two words, under the SC Algorithm :D
//
func CalculateSCSW(__word1, __word2 string) float64 {

	_length1 := len(__word1)
	_length2 := len(__word2)
	_minLength := Min(_length1, _length2)

	// Special case (min length == 3)
	if (_minLength <= 3 && _length1 != _length2) || _length1 < _length2 {
		return 1.0 // To risc to assume that they are equivalents
	} else {
		// Commun case
		_distance := 0
		_diff := false
		for _idx := 0; _idx < _minLength; _idx++ {
			if __word1[_idx] != __word2[_idx] {
				_distance = _idx
				_diff = true
				break
			}
		}
		if _diff {
			return ((float64(_minLength - _distance)) / (float64(_minLength)))
		} else {
			return 0.0
		}
	}

}
