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

// S175
const (
	null_coords = iota
	mull_pen
	dep_head
	zero_val
	temp_val
	inf_val
	bad_vardef
	lo_mem_stat_max
	sentinel        = mem_top
	temp_head       = mem_top - 1
	hold_head       = mem_top - 2
	hi_mem_stat_min = mem_top - 2
	end_attr        = temp_val
)


// s176
func initialize_table_entries () {
    var_used = integer(len(mem)+1)
    for len(int_name)<=max_given_internal {
        int_name = append(int_name, 0)
    }
    initialize_table_entries_A()
}