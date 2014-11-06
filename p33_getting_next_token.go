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

// s661
func check_outer_validity() bool {
	if scanner_status == normal {
		return true
	}
	jump_out(errors.New("check_outer_validity not implemented"))
	return false
}

// s665
func runaway() {
	if scanner_status > flushing {
		print_nl("Runaway ")
		switch scanner_status {
		case absorbing:
			print("text")
		case var_defining, op_defining:
			print("definition")
		case loop_defining:
			print("loop?")
		}
		print_ln()
		show_token_list(getLink(hold_head), null, error_line-10, 0)
	}
}

// s667
func get_next() {
	//mterror("get_next not implemented yet!")
	//jump_out(errors.New("get_next not implemented yet!"))
	var k halfword
	//var n integer
restart:
	cur_sym = 0
	if file_state() {
		// s 669 input from external file
		//jump_out(errors.New("s669 not implemented yet"))
	switch_label:
		fmt.Println("buffer[:20]", buffer[:20])
		fmt.Println("cur_input.loc", cur_input.(*inStateFileRec).loc)
		c := buffer[cur_input.(*inStateFileRec).loc]
		fmt.Println("c from buffer", string(c))
		cur_input.(*inStateFileRec).loc++
		class := char_class[c]
		fmt.Println("char and class in get_next s667:", c, class)
		switch class {
		case digit_class:
			goto start_numeric_token
		case period_class:
			class = char_class[buffer[cur_input.(*inStateFileRec).loc]]
			if class > period_class {
				goto switch_label
			} else {
				if class < period_class {
					class = digit_class
				}
				//n = 0  // REDO this when start_decimal implemented
				goto start_decimal_token
			}
		case space_class:
			goto switch_label
		case percent_class:
			// s679 move to next line of file or goto restart if there is
			// no next line
			if cur_input.getName() > 2 {
				// s681 read next line of file into buffer or goto restart
				// if file has ended
				jump_out(errors.New("s681 not implemented"))
			} else {
				if input_ptr > 0 {
					end_file_reading()
					goto restart
				}
				if selector < log_only {
					open_log_file()
				}
				if interaction > nonstop_mode {
					if cur_input.(*inStateFileRec).limit == cur_input.(*inStateFileRec).start {
						print_nl("(Please type a command or say 'end')")
					}
					print_ln()
					first = cur_input.(*inStateFileRec).start
					fmt.Println("calling prompt_input")
					prompt_input("*")
					fmt.Println("back from prompt_input")
					cur_input.(*inStateFileRec).limit = last
					buffer[cur_input.(*inStateFileRec).limit] = '%'
					first = cur_input.(*inStateFileRec).limit + 1
					cur_input.(*inStateFileRec).loc = cur_input.(*inStateFileRec).start
		            fmt.Println("buffer[:10]", buffer[:20])
		            fmt.Println("cur_input", cur_input)
				} else {
					fatal_error("*** (job aborted, no legal end found)")
				}
			}
			check_interrupt()
			goto switch_label
		case string_class:
			jump_out(errors.New("s671 not implemented"))
		case 5, 6, 7, 8: //isolated_classes
			k = cur_input.(*inStateFileRec).loc - 1
			goto found
		case invalid_class:
			fmt.Println("case invalid_class", string(rune(c)), class)
			print_err("Test line contains an invalid character")
			help("A funny symbol that I can't read has just been input.",
				"Continue, and I'll forget that it ever happened.")
			deletions_allowed = false
			mterror()
			deletions_allowed = true
			goto restart
			//fmt.Println("char_class['/']", char_class['/'])
			//fmt.Println("char_class['\\']", char_class['\\'])
			//jump_out(errors.New("s670 not implemented"))
		default:
			fmt.Println("do_nothing() for class", class)
			do_nothing()
		}
		fmt.Println("in get_next()", cur_input, cur_input)
		k = cur_input.(*inStateFileRec).loc - 1
		for ; char_class[buffer[cur_input.(*inStateFileRec).loc]] == class; cur_input.(*inStateFileRec).loc++ {
			//fmt.Println("counting over chars in class", class)
		}
		goto found
	start_numeric_token:
		jump_out(errors.New("s673 not implemented"))
	start_decimal_token:
		jump_out(errors.New("s674 not implemented"))
		//fin_numeric_token:
		jump_out(errors.New("s675 not implemented"))
	found:
		//cur_sym = id_lookup(k, cur_input.loc-k)
		fmt.Println("found: ", k, cur_input.(*inStateFileRec).loc, buffer[k:cur_input.(*inStateFileRec).loc])
		cur_sym = halfword(string_to_pos[string(buffer[k:cur_input.(*inStateFileRec).loc])])
		fmt.Println("cur_sym", cur_sym)
	} else {
		// s 676 input from token list
		jump_out(errors.New("s676 not implemented"))
	}
	// s668
	fmt.Printf("made it to s668 in get_next(). eqtb[cur_sym]:%T: %s\n", eqtb[cur_sym], eqtb[cur_sym])
	cur_cmd = eqtb[cur_sym].eq_type
	cur_mod = integer(eqtb[cur_sym].equiv)
	fmt.Println("s668 cur_cmd", cur_cmd, "cur_mod", cur_mod)
	if cur_cmd >= outer_tag {
		if check_outer_validity() {
			cur_cmd = cur_cmd - outer_tag
		} else {
			goto restart
		}
	}
}
