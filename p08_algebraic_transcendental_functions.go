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

// s 121
func square_rt(x scaled) scaled {
	var y int
	if x <= 0 {
		print_err("Square root of ")
		x.Print()
		print(" has been replaced by 0")
		help("Since I don't take square toots of negative numbers",
			"I'm zeroing this one., Proceed, with fingers crossed.")
		mterror()
		return 0
	}
	k := 23
	q := 2
	for x < fraction_two {
		k--
		x = x + x + x + x
	}
	if x < fraction_four {
		y = 0
	} else {
		x -= fraction_four
		y = 1
	}
	for {
		x += x
		y += y
		if x >= fraction_four {
			x -= fraction_four
			y++
		}
		x += x
		y = y + y - q
		q += q
		if x >= fraction_four {
			x -= fraction_four
			y++
		}
		if y > q {
			y -= q
			q += 2
		} else {
			if y <= 0 {
				q -= 2
				y += q
			}
		}
		k--
		if k == 0 {
			break
		}
	}
	return scaled(q >> 1)
}

// s150
func init_randoms(seed scaled) {
	fmt.Println("init_randoms not implemented")
}
