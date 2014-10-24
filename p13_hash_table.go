//	Copyright 2014 Thomas B. Hickey (thomasbhickey@gmail.com)
//
//	Licensed under the Apache License, Version 2.0 (the "License");
//	you may not use this file except in compliance with the License.
//	You may obtain a copy of the License at
//
//		http://www.apache.org/licenses/LICENSE-2.0
//
//	Unless required by applicable law or agreed to in writing, software
//	distributed under the License is distributed on an "AS IS" BASIS,
//	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//	See the License for the specific language governing permissions and
//	limitations under the License.
//
// MetaTrue is to a great extent a reimplementation of Donald Knuth's MetaFont

package metatrue

import (
    //"fmt"
)

var eqtb [] two_halves

// s210
func primitive(rs string, c, o halfword) {
    cur_sym = halfword(make_string(rs))
    for int(cur_sym)>=len(eqtb) {
        eqtb = append(eqtb, two_halves{lh:0, rh:0})
    }
    eqtb[cur_sym] = two_halves {lh:c, rh:o}
}
