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


// s796, s797

var cur_exp integer
var cur_type small_number

// s808
func flush_cur_exp(v scaled) {
    fatal_error("flush_cur_exp not implemented")
// 	switch cur_type {
// 	case unknown_types,
// 		transform_type,
// 		pair_type,
// 		dependent,
// 		proto_dependent,
// 		independent:
// 		recycle_value(cur_exp)
// 		free_node(cur_exp, value_node_size)
// 	case pen_type:
// 		delete_pen_ref(cur_exp)
// 	case string_type:
// 		delete_str_ref(cur_exp)
// 	case future_pen, path_type:
// 		toss_knot_list(cur_exp)
// 	case picture_type:
// 		toss_edges(cur_exp)
// 	}
// 	cur_type = known
// 	cur_exp = integer(v)
}

// s809
func recycle_value(integer) {
}

// s820
func flush_error(v scaled) {
	mterror()
	flush_cur_exp(v)
}

func put_get_error() {
    fatal_error("put_get_error not implemented")
	//back_error()
	//get_x_next()
}

func put_get_flush_error(v scaled) {
    fatal_error("put_get_flush not implemented")
	//put_get_error()
	//flush_cur_exp(v)
}
