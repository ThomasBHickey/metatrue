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

// s324
const (
	zero_w = 4
	void   = null + 1
)

// s327
var cur_edges pointer
var cur_wt integer

// s 332
func print_edges_sn(sn str_number, nuline bool, x_off, y_off integer) {
	print_edges(pos_to_string[sn], nuline, x_off, y_off)
}

func print_edges(s string, nuline bool, x_off, y_off integer) {
	fatal_error("print_edges not defined " + s)
}
