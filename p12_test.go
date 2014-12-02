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
	"testing"
)

func TestS189(t *testing.T) {
	if pen_offset_of != 100 {
		t.Error("S189 constants bad", pen_offset_of, y_scaled)
	}
}

func TestS190(t *testing.T) {
	set_up_internals()
	if max_given_internal != 41 {
		fmt.Println("S190 constants bad", max_given_internal)
	}
}

func TestS191(t *testing.T) {
	initialize() // should call set_up_internals
	if len(internal) != max_given_internal+1 {
		t.Error("S191 internal length wrong", len(internal), max_given_internal)
	}
	putPrimitivesIn_s192()
	rs := "boundarychar"
	fmt.Println("primitive", rs, make_string(rs), pos_to_string[make_string(rs)])
	fmt.Println("eqtb[296]", eqtb[296])
}

func TestS198(t *testing.T) {
	if string_class != 4 {
		t.Error("right_paren_class", right_paren_class)
	}
}

func TestS199(t *testing.T) {
	fmt.Println("char_class['\\']", char_class['\\'])
}
