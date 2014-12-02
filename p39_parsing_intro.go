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
	"fmt"
)

// s796, s797

var cur_exp integer
var cur_type small_number

// s799
func stash_cur_exp() pointer {
	var p pointer
	switch cur_type {
	case unknown_boolean, unknown_string, unknown_pen, unknown_picture, unknown_path,
		transform_type, pair_type, dependent, proto_dependent, independent:
		p = pointer(cur_exp)
	default:
		fmt.Println("name_type", capsule, "kind", cur_type, "value", cur_exp)
		p = get_avail(&value_tok{name_type: capsule, kind: cur_type, value: cur_exp, link: 133})
		// 		p = get_node(value_node_size)
		// 		name_type(p) = capsule
		// 		Type(p) = cur_type
		// 		value(p) = cur_exp
	}
	cur_type = vacuous
	//link(p) = void
	fmt.Printf("About to set link to void at %d: %T, %#v\n", p, mem[p], mem[p])
	//mem[p].(value_tok).link = void
	mem[p].setLink(void)
	fmt.Println("mem[p].(value_tok).getLink()", mem[p].(*value_tok).link)
	//var value_tok = mem[p]
	//fmt.Println("value_tok", value_tok)
	fmt.Println("mem[p].getLink()", mem[p].getLink())
	fmt.Printf("Node is now %#v\n", mem[p])
	//fmt.Println("Skippig void assignment in p39, s799")
	//mem[p].(num_tok).setLink(halfword(void))
	return p
}

// s800
func unstash_cur_exp(p pointer) {
	cur_type = getType(p)
	fmt.Println("unstash_cur_exp s802, cur_type:", cur_type)
	switch cur_type {
	case unknown_boolean, unknown_string, unknown_pen, unknown_picture, unknown_path,
		transform_type, pair_type, dependent, proto_dependent, independent:
		cur_exp = integer(p)
	default:
		fmt.Println("unstash_cur_exp default case:", p)
		cur_exp = integer(getValue(p))
		free_node(p)
	}
}

// s801

func print_exp(p pointer, verbosity small_number) {
	var (
		restore_cur_exp bool
		t               small_number
		v               integer
		//q pointer  // not needed until s803 Display a big node implemented
	)
	if p != null {
		restore_cur_exp = false
	} else {
		p = stash_cur_exp()
		restore_cur_exp = true
	}
	t = mem[p].getType()
	if t < dependent {
		v = integer(getValue(p))
	} else {
		if t < independent {
			v = integer(dep_list(p))
		}
	}
	// s802  // Print an abbreviated value of v with format depending on t
	switch t {
	case vacuous:
		print("vacuous")
	case boolean_type:
		if v == true_code {
			print("true")
		} else {
			print("false")
		}
	case unknown_boolean, unknown_string, unknown_pen, unknown_picture, unknown_path, numeric_type:
		// s806 Display a variable that's been declared but not defined
		print_type(t)
		if v != null {
			print_char(' ')
			for (getName_type(pointer(v)) == capsule) || (pointer(v) != p) {
				v = integer(getValue(pointer(v)))
			}
			print_variable_name(pointer(v))
		}
	case string_type:
		print_char('"')
		slow_print_sn(str_number(v))
		print_char('"')
	case pen_type, future_pen, path_type, picture_type:
		// s804 Display a complex type
		if verbosity <= 1 {
			print_type(t)
		} else {
			if selector == term_and_log {
				if internal[tracing_online] <= 0 {
					selector = term_only
					print_type(t)
					print(" (see the transcript file)")
					selector = term_and_log
				}
			}
			switch t {
			case pen_type:
				print_pen(pointer(v), "", false)
			case future_pen:
				print_path(pointer(v), " (future_pen)", false)
			case path_type:
				print_path(pointer(v), "", false)
			case picture_type:
				cur_edges = pointer(v)
				print_edges("", false, 0, 0)
			}
		}
	case transform_type, pair_type:
		if v == null {
			print_type(t)
		} else {
			// s803 Display a big node
			print_char('(')
			print("s803 Display a big node not implemented")
			print_char(')')
		}
	case known:
		scaled(v).Print() // print_scaled(v)
	case dependent, proto_dependent:
		print_dp(t, pointer(v), verbosity)
	case independent:
		print_variable_name(p)
	default:
		confusion("exp")
	}
	if restore_cur_exp {
		unstash_cur_exp(p)
	}
}

// s805
func print_dp(t small_number, p pointer, verbosity small_number) {
	fatal_error("print_dp not implemented")
}

// s807
func exp_err_sn(sn str_number) {
	disp_err_sn(null, sn)
}

func exp_err(s string) {
	disp_err(null, s)
}

func disp_err_sn(p pointer, sn str_number) {
	disp_err(p, pos_to_string[sn])
}

func disp_err(p pointer, s string) {
	if interaction == error_stop_mode {
		wake_up_terminal()
	}
	print_nl(">> ")
	print_exp(p, 1)
	if s != "" {
		print_nl("! ")
		print(s)
	}
}

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
func recycle_value(pointer) {
	fatal_error("recycle_value not implemented")
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

// s821
var var_flag small_number
