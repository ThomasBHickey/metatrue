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
	"math/big"
)

// s95
const el_gordo = 017777777777 // 2^31 -1

// s96
// we'll just code half(#) inline ourselves

// s97, s98
var arith_error bool = false

// s99
func chekc_arith() {
	if arith_error {
		clear_arith()
	}
}

func clear_arith() {
	print_err("Arithmetic overflow")
	help("Uh, oh. A little while ago one of the quantities that I was",
		"computing got too large, so I'm afraid your answers will be",
		"somewhat askew. You'll probably have to adopt different",
		"tactics next time. But I shall try to carry on anyway.")
	mterror()
	arith_error = false
}

// s100 Since we're going with Go's rational package, shouldn't need slow_add

// s101
const (
	quater_unit        = 040000  // 2^14 is 0.25
	half_unit          = 0100000 // 2^15 is 0.5
	three_quarter_unit = 0140000 // 3*2^14 is 0.75
	unity              = 0200000 // 2^16 is 1.0
	two                = 0400000 // 2^17 is 2.0
	three              = 0600000 // 2^17 + 2^16 is 3.0
)

// s102
type scaled big.Rat

func round_decimals(k int) *scaled {
	var a int64
	for k > 0 {
		k--
		a = (a + int64(dig[k])*two) / 10
	}
	return (*scaled)(big.NewRat((a+1)>>1, unity))
}
