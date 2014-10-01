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

// s989
func do_statement(){
    fatal_error("do_statement in 43/s989 not implemented")
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