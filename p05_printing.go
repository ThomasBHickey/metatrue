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
    "bufio"
	"fmt"
	//"os"
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
	log_file    *bufio.Writer
	selector    int = term_only
	dig         [23]int
	tally       int = 0 // s55
	term_offset int = 0
	file_offset int = 0
	trick_buf   [error_line + 1]rune
	trick_count int
	first_count int
)

// s55
// see s61 below

// s56
func wterm(msg string) {
	write(term_out, msg)
}
func wterm_ln(msg string) {
	write_ln(term_out, msg)
}
func wterm_cr() {
	write_ln(term_out)
}
func wlog(msg string) {
    //fmt.Println("in wlog", msg)
	write(log_file, msg)
}
func wlog_ln(msg string) {
	write_ln(log_file, msg)
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
	case no_print, pseudo, new_string:
	}
}

// s58
func print_char(r rune) {
	sr := string(r)
	//fmt.Println("print_char", sr)
	//fmt.Println("selector", selector)
	switch selector {
	case term_and_log:
		wterm(sr)
		wlog(sr)
		term_offset++
		file_offset++
		if term_offset == max_print_line {
			wterm_cr()
			term_offset = 0
		}
		if file_offset == max_print_line {
			wlog_cr()
			file_offset = 0
		}
	case log_only:
	    //fmt.Println("log_only")
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
	case no_print:
	case pseudo:
		if tally < trick_count {
			trick_buf[tally%error_line] = r
		}
	case new_string:
	}
	tally++
}

// s59
func print(s string) {
    //fmt.Println("in print()", s)
	for _, r := range s {
		print_char(r)
	}
}

func print_sn(sn str_number){
    print(pos_to_string[sn])
}

// s60
func slow_print(s string) {
	print(s)
}

func slow_print_sn(sn str_number){
    slow_print(pos_to_string[sn])
}

// s61
func initialize_output_routines() {
	wterm(banner)
	if base_ident == "" {
		wterm_ln(" (no base preloaded)")
	}
	update_terminal()
}

// s62
// Evidently a native Pascal operation
func odd(i int) bool {
	return (i % 2) == 1
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func print_nl(s string) {
	if ((term_offset > 0) && (odd(selector))) || ((file_offset > 0) && (selector >= log_only)) {
		print_ln()
	}
	print(s)
}

func print_nl_sn(sn str_number) {
    print_nl(pos_to_string[sn])
}

// s63
func print_the_digs(k int) {
	for k > 0 {
		print_char(rune(fmt.Sprintf("%d", k)[0]))
	}
}

// s64
// not sure this will pass Trap, but worth a try!
func print_int(n integer) {
	print(fmt.Sprintf("%d", n))
}

// s65
func print_dd(n integer) {
	print_int(integer(Abs(int(n)) % 100))
}

// s66
func prompt_input(s string) {
    wake_up_terminal()
    print(s)
    term_input()
}

func term_input() {
    update_terminal()
    if ! input_ln(term_in, true) {
        fatal_error("end of file on the terminal!")
    }
    term_offset = 0
    selector--
    if last != first{
        print(string(buffer[first:last]))
    }
    print_ln()
    buffer[last] = '%'
    selector++
}
