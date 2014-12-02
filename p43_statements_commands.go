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

// s989
func do_statement() {
	fmt.Println("in do_statement in 43/s989")
	cur_type = vacuous
	get_x_next()
	if cur_cmd > max_statement_command {
		// do s993 equation, assignment, title or <expression> endgroup
		var_flag = assignment
		scan_expression()
		if cur_cmd == equals {
			do_equation()
		} else {
			if cur_cmd == assignment {
				do_assignment()
			} else {
				if cur_type == string_type {
					// do a title 994
					if internal[tracing_titles] > 0 {
						print_nl("")
						slow_print_sn(str_number(cur_exp))
						update_terminal()
					}
					if internal[proofing] > 0 {
						// send current expression as a title to the output file
						jump_out(errors.New("1179 not implemented"))
					}
				} else {
					if cur_type != vacuous {
						exp_err("Isolated expression")
						help("I couldn't find an '=' or ':=' after the",
							"expression that is shown abovve this error message,",
							"so I guess I'll just ignore it and carry on.")
						put_get_error()
					}
				}
			}
		}
		flush_cur_exp(0)
		cur_type = vacuous
	} else {
		// do s992 a statement that doesn't begin with an expression
		if internal[tracing_commands] > 0 {
			show_cur_cmd_mod()
		}
		switch cur_cmd {
		case type_name:
			do_type_declaration()
		case macro_def:
			if cur_mod > var_def {
				make_op_def()
			} else {
				if cur_mod > end_def {
					scan_def()
				}
			}
		// s1020 cases of do_statement for particular commands
		default:
			jump_out(errors.New("s1020 not implemented"))
		}
	}
	if cur_cmd < semicolon {
		// s991 Flush unparable junk that was found after the statment
	}
	error_count = 0
	fmt.Println("finished do_statement")
}

// s995
func do_equation() {
	jump_out(errors.New("do_equation not implemented"))
}

// s996
func do_assignment() {
	jump_out(errors.New("do_assignment not implemented"))
}

// s1017
func main_control() {
	for {
		do_statement()
		if cur_cmd == end_group {
			print_err("Extra `endgroup'")
			help("I'm not currently working on a `begingroup',",
				"so I had better not try to end anything.")
			flush_error(0)
		}
		if cur_cmd == stop {
			break
		}
	}
}

// s1015
func do_type_declaration() {
	jump_out(errors.New("do_type_declaration not implemented"))
}
