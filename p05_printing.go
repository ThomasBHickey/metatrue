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

// s54
const (
	no_print = iota
	term_only
	log_only
	term_and_log
	pseudo
	new_string
	max_selector = new_string
)

var (
	log_file    *os.File
	selector    int = term_only
	dig         [23]int
	tally       int = 0 // s55
	term_offset int = 0
	file_offset int = 0
	trick_buf   [error_line + 1]rune
	trick_count int
	first_count int
)

// s56
func wterm(msgs ...string) {
	write(term_out, msgs)
}
func wterm_ln(msgs ...string) {
	write_ln(term_out, msgs)
}
func wterm_cr() {
	write_ln(term_out)
}
func wlog(msgs ...string) {
	write(log_file, msgs)
}
func wlog_ln(msgs ...string) {
	write_ln(log_file, msgs)
}
func wlog_cr() {
	write_ln(log_file)
}

// s57




