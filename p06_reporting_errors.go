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
	"os"
	"unicode"
)

var help_string string

// s68

const (
	batch_mode = iota
	nonstop_mode
	scroll_mode
	error_stop_mode
)

func print_err(msg string) {
	if interaction == error_stop_mode {
		wake_up_terminal()
		print_nl("! ")
		print(msg)
	}
}

// s 68, 69
var interaction int = error_stop_mode

// s70
func initialize_print_selector() {
	if interaction == batch_mode {
		selector = no_print
	} else {
		selector = term_only
	}
}

// s71, 72
const (
	spotless = iota
	warning_issued
	error_message_issued
	fatal_error_stop
)

var deletions_allowed = true
var history int
var error_count int = 0

// s74
var use_err_help = false
var err_help = ""
var help_lines = []string{}

func help(msgs ...string) {
	//for _, msg := range msgs {
	//	print(msg)
	//}
	help_lines = msgs
}

// s 76
func jump_out(err error) {
	if err != nil {
		fmt.Println("MF Terminating:", err)
	}
	close_files_and_terminate()
	os.Exit(1)
}

// s77
func mterror() {
	if history < error_message_issued {
		history = error_message_issued
	}
	print_char('.')
	show_context()
	if interaction == error_stop_mode {
		get_users_advice() // in s78
		return
	}
	error_count++
	if error_count == 100 {
		print_nl("(That makes 100 errors; please try again.)")
		history = fatal_error_stop
		jump_out(nil)
	}
	put_help_msg_on_transcript()
}

// s78
func get_users_advice() {
	for {
		clear_for_error_prompt()
		prompt_input("? ")
		if last == first {
			return
		}
		c := unicode.ToUpper(buffer[first])
		switch { // in s79
		case '0' <= c && c <= '9':
			delete_tokens(c)
			continue
		case 'D' == c && debug:
			debug_help()
			continue
		case 'E' == c && file_ptr > 0:
			print_nl("You want to edit file ")
			slow_print_sn(input_stack[file_ptr].getName())
			print(" at line ")
			print_int(line)
			interaction = scroll_mode
			jump_out(nil)
		case 'H' == c:
			print_help_info()
			continue
		case 'I' == c:
			introduce_new_material() // in 82
			return
		case 'Q' == c || 'R' == c || 'S' == c:
			change_interaction_level(c) // in 81
			return
		case 'X' == c:
			interaction = scroll_mode
			jump_out(nil)
		}
		print_menu_of_available_options() // s80
	}
}

// s80
func print_menu_of_available_options() {
	print("Type <return> to proceeed, S to scroll future error _messages,")
	print_nl("R to run without stopping, Q to run quietly,")
	print_nl("I to insert something, ")
	if file_ptr > 0 {
		print("E to edit your file,")
	}
	if deletions_allowed {
		print_nl("! or ... or 9 to ignore the next 1 to 9 tokens of input,")
	}
	print_nl("H for help, X to quit.")
}

// s81
func change_interaction_level(c rune) {
	error_count = 0
	interaction = batch_mode + int(c-'Q')
	print("OK, entering")
	switch c {
	case 'Q':
		print("batchmode")
		selector--
	case 'R':
		print("nonstopmode")
	case 'S':
		print("scrollmode")
	}
	print("...")
	print_ln()
	update_terminal()
}

// s82
func introduce_new_material() {
	if last > first+1 {
		cur_input.(*inStateFileRec).loc = first + 1
		buffer[first] = ' '
	} else {
		prompt_input("insert>")
		cur_input.(*inStateFileRec).loc = first
	}
	first = last + 1
	cur_input.(*inStateFileRec).limit = last
}

// s83
func delete_tokens(c rune) {
	var (
		s1 = cur_cmd
		s2 = cur_mod
		s3 = cur_sym
	)
	OK_to_interrupt = false
	if (last > first+1) && (buffer[first+1] >= '0') && (buffer[first+1] <= '9') {
		c = c*10 + buffer[first+1] - '0'*11
	} else {
		c = c - '0'
	}
	for c > 0 {
		get_next()
		c--
	}
	cur_cmd = s1
	cur_mod = s2
	cur_sym = s3
	OK_to_interrupt = true
	help("I have just deleted som text, as you asked.",
		"You can now delete more, or insert, or whatever.")
	show_context()
}

// s84

func print_help_info() {
	if use_err_help {
		// s85
		print(err_help)
		use_err_help = false
	} else {
		if len(err_help) == 0 {
			help("Sorry, I don't know how to help in this situation.",
				"Maybe you should try asking a human?")
		}
	}
	err_help = "Sorry, I already gave what help I could...]n" +
		"Maybe you shold try asking a human?\n" +
		"An error might have occurred before I noticed any problems.\n" +
		"''If all else fails, read the instrucitons.''\n"
}

// s86
func put_help_msg_on_transcript() {
	if interaction > batch_mode {
		selector--
	}
	if use_err_help {
		print_nl("")
		print(err_help)
	} else {
		print(help_string)
	}
	if interaction > batch_mode {
		selector++
	}
	print_ln()
}

// s87
func normalize_selector() {
	if log_opened {
		selector = term_and_log
	} else {
		selector = term_only
	}
	if job_name == "" {
		open_log_file()
	}
	if interaction == batch_mode {
		selector--
	}
}

// s88
func succumb() {
	fmt.Println("In succumb()")
	if interaction == error_stop_mode {
		interaction = scroll_mode
	}
	if log_opened {
		mterror()
	}
	if debug {
		if interaction > batch_mode {
			debug_help()
		}
	}
	history = fatal_error_stop
	jump_out(nil)
}

func fatal_error(s string) {
	//panic("in fatal error (needs fixing): "+s)
	//jump_out(nil)
	fmt.Println("fatal_error " + s)
	normalize_selector()
	print_err("Emergency stop")
	help(s)
	succumb()
}

// s89
func overflow(errmsg string, n int) {
	normalize_selector()
	print_err(fmt.Sprintf("METATRUE capacity exceeded, sorry [%s=%d]", errmsg, n))
	help("If you really absolutely need more capacity",
		"you can ask a wizard to enlarge me.")
	succumb()
}

// s90
func confusion(s string) {
	normalize_selector()
	if history < error_message_issued {
		print_err("This can't happen(")
		print(s)
		print_char(')')
		help("I'm broken. Please show this to someone who can fix can fix")
	} else {
		print_err("I can't go on meeting you like this")
		help_string = "One of your faux pas seems to have wounded me deeply...\n" +
			"in fact, I'm barely conscious. Please fix it and try again."
	}
	succumb()
}

// s91
func check_interrupt() {
	if interrupt != 0 {
		pause_for_instructions()
	}
}

// s 92
var (
	interrupt       = 0
	OK_to_interrupt = true
)

// s93
func pause_for_instructions() {
	if OK_to_interrupt {
		interaction = error_stop_mode
		if (selector == log_only) || (selector == no_print) {
			selector++
		}
		print_err("Interruption")
		deletions_allowed = false
		help("You rang?\n",
			"Try to insert some instructions for me (e.g. 'I show x'),",
			"unless you just want to quit by typing 'X'.")
		mterror()
		deletions_allowed = true
		interrupt = 0
	}
}

// s94
func missing_err(s string) {
	print_err("missing '")
	print(s)
	print("' has been inserted")
}
