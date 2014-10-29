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
	"errors"
	"fmt"
)

// s214
// Pascal MetaFont stores these in the one-word section of mem
type sym_tok struct {
	p halfword
}

func (node *sym_tok) Type() small_number {
	return symbolic
}

type num_tok struct {
	value     scaled
	link      pointer
	info      halfword
	name_type quarterword
}

func (node *num_tok) Type() small_number {
	return known
}

func (node *num_tok) setLink(l pointer){
	node.link = l
}

// func name_type(p pointer) quarterword {
// 	return mem[p].(num_tok).name_type
// }

type string_tok struct {
	value str_number
}

func (node *string_tok) Type() small_number {
	return string_type
}

func (node *string_tok) setLink(p pointer){
	jump_out(errors.New("Tried to call setLink of a string_tok!"))
}

type value_tok struct {
	kind      small_number
	name_type small_number
	link      pointer
	value     integer
}

func (node *value_tok) Type() small_number {
	return node.kind
}

func (node *value_tok) setLink(link pointer){
	fmt.Println("setting link in value_tok", link, node)
	node.link = link
	fmt.Println("value after setting link", node)
}

func (node *value_tok) getLink() pointer {
	return node.link
}

func value_loc(p pointer) pointer {
	return p + 1
}

// func value(p pointer) integer {
// 	return integer(mem[p].(num_tok).value)
// }

var expr_base, suffix_base, text_base halfword

func setup_expr_base() {
	expr_base = halfword(len(pos_to_string))
	for i := 0; i < param_size; i++ {
		_ = append(pos_to_string, "xxxx")
	}
	suffix_base = halfword(len(pos_to_string))
	for i := 0; i < param_size; i++ {
		_ = append(pos_to_string, "yyyy")
	}
	text_base = halfword(len(pos_to_string))
	for i := 0; i < param_size; i++ {
		_ = append(pos_to_string, "zzzz")
	}
}

// s215
func new_num_tok(v scaled) pointer {
	node := &num_tok{value: v}
	return get_avail(node)
}

func Type(p pointer) small_number {
	return mem[p].Type()
}

// s216
func flush_token_list(p pointer) {
	var q pointer
	for p != null {
		q = p
		p = link(p)
		switch Type(q) {
		case vacuous, boolean_type, known:
		case string_type:
			delete_str_ref(mem[q].(*string_tok).value)
		case unknown_boolean, unknown_string, unknown_pen, unknown_picture, unknown_path,
			pen_type, path_type, future_pen, picture_type, transform_type, dependent,
			proto_dependent, independent:
			g_pointer = q
			token_recycle()
		default:
			confusion("token")
		}
		free_node(q)
	}
}

// S217
func show_token_list(p pointer, q integer, l, null_tally integer) {
	fatal_error("show_token_list not implemented")
}

// s224
func print_capsule() {
	print_char('(')
	print_exp(g_pointer, 0)
	print_char(')')
}

func token_recycle() {
	recycle_value(g_pointer)
}

// s225
var g_pointer pointer
