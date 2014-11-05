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
	"bufio"
	"fmt"
)

// s627
// dropped the trailing "_field" for each field
type InputStateRec interface {
    getIndex() quarterword
    getName() str_number
}

type inStateFileRec struct {
	index quarterword
	start,
	loc,
	limit halfword
	name str_number
}

func (input_state inStateFileRec) getIndex() quarterword {
    return input_state.index
}

func (input_state inStateFileRec) getName() str_number {
    return input_state.name
}

// s628
var (
	input_stack  [stack_size + 1]InputStateRec
	input_ptr    = 0
	max_in_stack = 0
	cur_input    InputStateRec
)

// s631
var (
	in_open     quarterword
	open_parens integer
	input_file  [max_in_open + 1]*bufio.Reader
	line        integer
	line_stack  [max_in_open + 1]integer
)

func terminal_input() bool {
	return cur_input.(*inStateFileRec).name == make_string("")
}

func cur_file() *bufio.Reader {
	return input_file[cur_input.(*inStateFileRec).index]
}

// s 632
const (
    forever_text = iota+max_in_open
    loop_text
    parameter
    backed_up
    inserted
    macro
    )
type inStateListRec struct {
    token_type quarterword
    start pointer  // first  node of the token list
    loc,    // poiinter to current node in token list
    param_start halfword
    name str_number
}

func (input_state inStateListRec) getName() str_number {
    return input_state.name
}


func (input_state *inStateListRec) getIndex() quarterword {
    return input_state.token_type
}

func file_state() bool {
    fmt.Println("file_state <=?", cur_input.getIndex(), max_in_open)
	return cur_input.getIndex() <= max_in_open
}

// s633
var (
	param_stack     [param_size + 1]pointer
	param_ptr       int
	max_param_stack int
)

// s634
var file_ptr int

// s635
func show_context() {
    fmt.Println("show_context not yet implemented")
    //var old_setting quarterword
    // s641 local variables
    var (
        //i halfword
        //l, m integer
        //n halfword
        //p, q integer
        )
    file_ptr = input_ptr
    input_stack[file_ptr] = cur_input
    for {
        cur_input = input_stack[file_ptr]
        // s 636 display the current context
        if (file_ptr==input_ptr)|| file_state() || (cur_input.(*inStateListRec).token_type!=backed_up) || (cur_input.(*inStateFileRec).loc!=null){
            tally = 0
            old_setting = selector
            if file_state() {
                //s 637 print location of current line
                fmt.Println("s637 not implemented yet")
                // s644 pseudoprint the line
                fmt.Println("s644 not implemented yet")
            } else {
                // s638 print type of token list
                // s 645 pseudoprint the token list
            }
                
        if file_state() {
            if (cur_input.getName()>2) || (file_ptr==0) { break}
        }
        file_ptr--
    }
    // done label
    cur_input = input_stack[input_ptr]
}
}
