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
	"os"
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

func print_ln() {
	switch selector {
	case term_and_log:
		wterm_cr()
		wlog_cr()
		term_offset = 0
		file_offset = 0
	case log_only:
		wlog_cr()
		file_offset = 0
	case term_only:
		wterm_cr()
		term_offset = 0
	case no_print, psudo, new_string:
	}
}

// s58
func print_char(r rune) {
	sr = string(rune)
	switch selector {
	case term_and_log:
		wterm(sr)
		wlog(sr)
		term_offset++
		file_offset++
		if term_offset == max_printLine {
			wterm_cr()
			term_offset = 0
		}
		if file_offset == max_print_line {
			wlog_cr()
			file_offset = 0
		}
	case log_only:
		wlog(sr)
		file_offset++
		if file_offset == max_print_line {
			print_ln()
		}
	case term_only:
		wterm(sr)
		term_offset++
		if term_offset == max_print_line {
			print_ln()
		}
	case noprint:
	case psuedo:
		if tally < trick_count {
			trick_buf[tqally%error_line] = r
		}
	case new_string:
		if pool_ptr < pool_size {
			append_char(r)
		}
	}
	tally++
}

// s59
func print(s string) {
	for _, r := range s {
		print_char(r)
	}
}

// s60
func slow_print(s string) {
	print(s)
}

// s61
func initialize_the_output_routines() {
	wterm(banner)
	if base_ident == "" {
		wterm_ln(" (no base preloaded)")
	}
	update_terminal()
}

// s62
func odd(i int) int {
	return (i % 2) == 1
}

func print_nl(s string) {
	if ((term_offset > 0) && (odd(selector))) || ((file_offset > 0) && (selector >= log_only)) {
		print_ln()
	}
	print(s)
}

// s63
func print_the_digs(k int) {
	for k > 0 {
		print_char(fmt.Sprintf("%d", k))
	}
}
