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

// s37...
type (
    pool_pointer pointer
    str_number halfword
)

var (
    string_to_pos = make(map[string] str_number)
    pos_to_string []string
)

// s43
func delete_str_ref(p str_number){
    // does nothing right now
}

// s44
func make_string(s string) str_number {
    pos, ok := string_to_pos[s]
    if ok {return pos}
    string_to_pos[s] = str_number(len(pos_to_string))
    pos_to_string = append(pos_to_string, s)
    return str_number(len(pos_to_string)-1)
}

// s47
func get_strings_started() {
    for r:= rune(0); r<rune(256); r++ {
        if character_cannot_be_printed(r){
            if r<0100 {
                make_string("^^"+string(r+0100))
            } else {
                if r<0200 { 
                    make_string("^^"+string(r-0100))
                }else {
                    make_string("^^"+fmt.Sprintf("%x", r))
                }
            }
        } else {
            make_string(string(r))
        }
    }
}

// s49
func character_cannot_be_printed(r rune) bool {
    return r<' ' || r>'~'
}
