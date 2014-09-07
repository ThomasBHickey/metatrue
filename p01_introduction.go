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

//s 7
const (
	debug = true
	stats = true
)

//s 11
const (
	buf_size        = 500 // max chars present in current lines of open files
	error_line      = 72
	half_error_line = 42
	max_print_line  = 79
	screen_width    = 768
	screen_depth    = 1024
	stack_size      = 30 //simultaneous input sources
	max_wiggle      = 300
	path_size       = 300
	bistack_size    = 785
	header_size     = 100
	lig_table_size  = 5000
	max_kerns = 500
	max_font_dimen = 50
)

// 12
const (
	max_in_open = 6
	param_size  = 150
)

// s13
var bad int

// s 14
// func checkConstantsForConsistency() error {
// 	switch true {
// 	case (half_error_line < 30) || (half_error_line > error_line-15):
// 		bad = 1
// 	case max_print_line < 60:
// 		bad = 2
// 	case header_size%4 != 0:
// 		bad = 6
// 	case (lig_table_size < 255) || (lig_table_size > 32510):
// 		bad = 7
// 	}

	// 	if (half_error_line < 30) || (half_error_line > error_line-15) {
	// 		bad = 1
	// 	}
	// 	if max_print_line < 60 {
	// 		bad = 2
	// 	}
	// 	if header_size%4 != 0 {
	// 		bad = 6
	// 	}
	// 	if (lig_table_size < 255) || (lig_table_size > 32510) {
	// 		bad = 7
	// 	}

//}
