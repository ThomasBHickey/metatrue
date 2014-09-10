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
)

// s68

const (
    batch_mode = iota
    nonstop_mode
    scroll_mode
    error_stop_mode
    )

func print_err(msg string){
    if interaction==error_stop_mode {
        wake_up_terminal()
        print_nl("! ")
        print(msg)
    }
}

var interaction int = error_stop_mode

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
func help(msgs ...string){
    for msg := range msgs {
        print(msg)
    }
}

// s 76
func jump_out(){
    close_files_and_terminate()
    os.Exit(1)
}


// s88
func succumb(){
    if interaction==error_stop_mode{
        interaction = scroll_mode}
    if log_opened { mterror()}
    if debug {
        if interaction > batch_mode {
            debug_help()
        }
    }
    history = fatal_error_stop;
}


// s89
func overflow(errmsg string, n int){
    normalize_selector()
    print_err(fmt.Sprintf("METATRUE capacity exceeded, sorry [%s=%d]", errmsg, n))
    help("If you really absolutely need more capacity",
    "you can ask a wizard to enlarge me.")
    succumb()
}
