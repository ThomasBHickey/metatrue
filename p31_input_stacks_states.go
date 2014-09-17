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

// s627
// dropped the trailing "_field" for each field
type in_state_record struct {
	index,
	start,
	loc,
	limit int
	name string
}

// s628
var (
	input_stack [stack_size + 1]in_state_record
	input_ptr    = 0
	max_in_stack = 0
	cur_input    in_state_record
)

// s631
func terminal_input() bool {
    return cur_input.name==""
}
func cur_file(){
    return input_file[index]
}

// s633
var (
    param_stack [param_size+1] int
    param_ptr int
    max_param_stack int
    )
    
// s634
var file_ptr int

// s635
func show_context(){
    print_err("showContext not yet implemented")
}