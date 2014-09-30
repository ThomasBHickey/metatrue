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

// s2
var banner = "This is METATRUE, Version 1"

// s 4
//Don't think we need to pull everthing together in Go

// s 6
// skip the labels

// s 7
// not macros, but should work fine
const (
	debug = true
	stat  = true
)

//s 8
// We're not going to do any preloading

//s 9, 10
// Pascal stuff handled outside the code in Go

//s 11
const (
	mem_max = 30000 // prevent runaways
	//max_internal    // internals are a map
	buf_size        = 500 // max chars present in current lines of open files
	error_line      = 72
	half_error_line = 42
	max_print_line  = 79
	screen_width    = 768
	screen_depth    = 1024
	stack_size      = 30 //simultaneous input sources
	//max_strings // skipped
	//string_vacancies
	//pool_size
	move_size  = 5000
	max_wiggle = 300
	//gf_buf_size  // depend on Go's buffering?
	file_name_size = 40 // might not need this
	//pool_name
	path_size      = 300
	bistack_size   = 785
	header_size    = 100
	lig_table_size = 5000
	max_kerns      = 500
	max_font_dimen = 50
)

// 12
const (
	max_in_open = 6
	param_size  = 150
	mem_min     = 0
	mem_top = 30000
)

// s13
var bad int

// s 14
// check constants in p01_Test.go

// s15
// try to do without labels for now

// s16 macros
// most of these we won't need since
// we don't have to worry about as many
// portability issues
