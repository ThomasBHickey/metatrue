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

// s652
func back_input() {
	fmt.Println("back_input not yet written")
}

// s655

// s656
func clear_for_error_prompt(){
    for file_state || terminal_input || (input_ptr>0) || (loc==limit){
        end_file_reading()
    }
    print_ln()
    clear_terminal()
}

// Section 657
// initializeTheInputRoutines is only called once in s1211, so global
// initializations (e.g. scanner_status) are fine

var (
	//input_ptr    = 0
	//max_in_stack = 0
	in_open = 0
	//max_buf_stack = 0
)

func initializeTheInputRoutines() error {
	fmt.Println("initializeTheInputRoutines")
	err := init_terminal()
	if err != nil {
		return err
	}
	cur_input.limit = last
	first = last + 1
	fmt.Println("initializeTheInputRoutines returning OK")
	return nil
}

// s659

const (
	normal = iota
	skipping
	flushing
	absorbing
	var_defining
	op_defining
	loop_defining
)

var (
	scanner_status = normal
	warning_info   int
)

// s660
//scanner_status = normal # set in s659