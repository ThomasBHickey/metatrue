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
// s214
type sym_tok struct{
    p halfword
}
func (node sym_tok) Type() quarterword{
    return symbolic
}
type num_tok struct {
    name_type quarterword
    value scaled
}
func (node num_tok) Type() quarterword{
    return known
}

// s215
func new_num_tok(v scaled) pointer {
    node := num_tok{value:v, name_type: token}
    return get_avail(node)
}

func Type(p pointer) quarterword{
    return mem[p].Type()
}

// s216
func flush_token_list(p pointer) {
    var q pointer
    for p!=null {
        q = p
        p = link(p)
        switch Type(q) {
            case vauous, boolean_type, known:
            case string_type: delete_str_ref(value(q))
            case unknown_boolean, unknown_string, unknown_pen, unknown_picture, unknown_path,
            pen_type, path_type, future_pen, picture_type, transform_type, dependent,
            proto_dependent, independent: g_pointer = q; token_recycle()
            default: confusion("token")
        }
        free_node(q)
    }
}
    
// S217
func show_token_list(p pointer, q integer, l, null_tally integer){
    fatal_error("show_token_list not implemented")
}
