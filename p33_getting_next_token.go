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
//"fmt"
)

// s665
func runaway() {
    if scanner_status > flushing {
        print_nl("Runaway ")
        switch scanner_status {
            case absorbing: print("text")
            case var_defining, op_defining: print("definition")
            case loop_defining: print("loop?")
        }
        print_ln()
        show_token_list(link(hold_head), null, error_line-10, 0)
    }
}

// s667
func get_next(){
    //mterror("get_next not implemented yet!")
    jump_out(errors.New("get_next not implemented yet!"))
}