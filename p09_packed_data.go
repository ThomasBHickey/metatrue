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

// s153

const (
	// if integer  is int32 bits:
	min_quarterword = 0
	max_quarterword = 255
	min_halfword    = 0
	max_halfword    = 65535

// if integer is int64 bits:
//min_quarterword = 0
//max_quaterword = 0xFFFF
//min_halfword = 0
//max_halfword = 0xFFFFFFFF
)

// s154 in p09_test.go

// s155

func ho(i int) int { return i }
func qo(i int) int { return i }
func qi(i int) int { return i }

// s156

type (
	// for integer int32
	quarterword byte
	halfword    uint16
	// if integer int64
	//quarterword uint16
	//halfword uint32
	//two_choices   byte
	//three_choices byte
	two_halves    struct {
		lh, rh halfword
	}
	half_two_quarters struct {
		lh     halfword
		b0, b1 quarterword
	}
	four_quarters struct {
		b0, b1, b2, b3 quarterword
	}
)
func (Node two_halves) Type() string {
    return "two_halves"
}

// s157
// I don't think we'll have things we don't know what type they are

// s158 and following:  we're going to try to use Go's memory allocation

// s175
// we'll need to define these fixed usage nodes at some point!
