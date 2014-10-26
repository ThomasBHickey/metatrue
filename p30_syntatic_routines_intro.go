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

// s624
var (
    cur_cmd halfword //eight_bits
	cur_mod integer
	cur_sym halfword
	)
	
// s625
func print_cmd_mod(c, m integer){
    switch c {
        // s 212 symbolic printing of primitives
        case 1: fmt.Println("s 212 not implemented")
    }
}

// s626
func show_cur_cmd_mod() {
    show_cmd_mod(integer(cur_cmd), cur_mod)
}

func show_cmd_mod(c, m integer){
    begin_diagnostic()
    print_nl("{")
    print_cmd_mod(c, m)
    print_char('}')
    end_diagnostic(false)
}
