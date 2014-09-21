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
	"bytes"
	"fmt"
	//"math/big"
	//"strings"
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

// s100 Originally planned on using Go's rational package.  But not too much
// magic in it, and it looks as though there is substantial overhead in
// using it.  int64 should avoid the normal overflow problems people run
// into with MetaFont and still make it possible to match MetaFont's results

// s101
const (
	quarter_unit       = 040000  // 2^14 is 0.25
	half_unit          = 0100000 // 2^15 is 0.5
	three_quarter_unit = 0140000 // 3*2^14 is 0.75
	unity              = 0200000 // 2^16 is 1.0
	two                = 0400000 // 2^17 is 2.0
	three              = 0600000 // 2^17 + 2^16 is 3.0
)

// s102
//type scaled big.Rat
type scaled int64

func round_decimals(k int) scaled {
	var a int64
	for k > 0 {
		k--
		a = (a + int64(dig[k])*two) / 10
		//a = (a + dig[k]*two) / 10
	}
	//return (*scaled)(big.NewRat((a+1)>>1, unity))
	return scaled((a + 1) / 2)
}

// s103

// added floatString for easier testing
// The big.Rat.FloatString routine returned "0.66667" for 2/3's
// this routine appears to return "0.66666"
func (s scaled) floatString() string {
	// 	ts := (*big.Rat)(sc).FloatString(5)
	// 	if strings.HasSuffix(ts, ".00000") {
	// 		return ts[:len(ts)-6]
	// 	}
	// 	return ts
	buffer := bytes.NewBuffer([]byte{})
	var delta scaled
	if s < 0 {
		buffer.WriteByte('-')
		//print_char('-')
		s = -s
	}
	fmt.Fprintf(buffer, "%d", s/unity)
	s = 10*(s%unity) + 5
	if s != 5 {
		delta = scaled(10)
		buffer.WriteByte('.')
		for {
			if delta > unity {
				s = s + 0100000 - (delta / 2)
			}
			buffer.WriteByte(byte('0' + (s / unity)))
			s = 10 * (s % unity)
			delta = delta * 10
			if s <= delta {
				break
			}
		}
	}
	return buffer.String()
}

func (s scaled) Print() {
	print(s.floatString())
}

// s104
func print_two(x, y scaled) {
	print_char('(')
	x.Print()
	print_char(',')
	y.Print()
	print_char(')')
}

// s105
const (
	fraction_half  = 01000000000
	fraction_one   = 02000000000
	fraction_two   = 04000000000
	fraction_three = 06000000000
	fraction_four  = 010000000000
)

type fraction int64

// s106
const (
	forty_five_deg  = 0264000000
	ninety_deg      = 0550000000
	one_eighty_deg  = 01320000000
	three_sixty_deg = 02640000000
)

type angle int64

// s107
// not worrying about overflow
func make_fraction(p, q int) fraction {
	return fraction((fraction_two*p + q) / (2 * q))
}

// s109
// still not worrying about overflow!
func take_fraction(q int64, f fraction) int64 {
	negative := f < 0
	if q < 0 {
		negative = !negative
	}
	if negative {
		f = -f
	}
	// 	fmt.Printf("take_fraction negative: %s\n", negative)
	// 	fmt.Printf("fraction_half %x\n", fraction_half)
	// 	fmt.Printf("fraction_one  %x\n", fraction_one)
	// 	fmt.Printf("f             %x\n", f)
	// 	fmt.Printf("q             %x\n", q)
	// 	fmt.Printf("f+fraction_half: %x\n", f+fraction_half)
	// 	fmt.Printf("fraction_one:    %x\n", fraction_one)
	// 	fmt.Printf("q*f+1/2): %x", q*int64(f)+fraction_half)
	p := (q*int64(f) + fraction_half) / fraction_one
	if negative {
		p = -p
	}
	return p
}

// s112
func take_scaled(q int64, f scaled) int64 {
	negative := f < 0
	if q < 0 {
		negative = !negative
	}
	if negative {
		f = -f
	}
	p := (q*int64(f) + half_unit) / unity
	if negative {
		p = -p
	}
	return p
}

// s114
// not worrying about overflow
func make_scaled(p, q int) scaled {
	return scaled((two*p + q) / (2 * q))
}
