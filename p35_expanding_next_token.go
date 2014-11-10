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

// s707
func expand(){
    var (
        //p pointer
        //k integer
        //j pool_pointer
    )
    if internal[tracing_commands]>unity {
        if cur_cmd!=defined_macro {
            show_cur_cmd_mod()
        }
    }
    fmt.Println("in expand, cur_cmd=", cur_cmd)
    switch cur_cmd {
        case if_test: conditional()
        case fi_or_else: // s751 terminate current conditional and skip to fi
                        fmt.Println("s751 not implemented")
        case input: //s 711 Initiate or terminate input from a file
        case iteration: if cur_mod==end_for {
            //s 708 Scold user for extra endfor
            fmt.Println("s708 not yet implemented")}else{
                begin_iteration()
            }
        case repeat_loop: //s 712 Repeat a loop
            fmt.Println("s 712 repeat_loop not implemented")
        case exit_test: // s713 Exit loop if the proper time has come
            fmt.Println("s 713 not implemented")
        case relax: 
            fmt.Println("found relax!")
            do_nothing()
        case expand_after: //s715 Expand token after next token
            fmt.Println("s 715 not implemented")
        case scan_tokens: // s716 put string into input buffer
            fmt.Println("s 716 not implemented")
        case defined_macro: macro_call(pointer(cur_mod), null, pointer(cur_sym))
    }
}

// s709
func putPrimitivesIn_s709() {
    primitive("input", input, 0)
    primitive("endinput", input, 1)
}

// s718
func get_x_next(){
    var save_exp pointer
    fmt.Println("calling get_next from get_x_next")
    get_next()
    fmt.Println("just called get_next in get_x_next, cur_cmd:", cur_cmd, "min_command", min_command)
    if cur_cmd<min_command {
        save_exp = stash_cur_exp()
        for cur_cmd<min_command {
            if cur_cmd==defined_macro {
                macro_call(pointer(cur_mod), null, pointer(cur_sym))
            } else { expand() }
            fmt.Println("2d call to get_next in get_x_next")
            get_next()
            fmt.Println("cur_cmd", cur_cmd)
        }
        unstash_cur_exp(save_exp)
    }
}


// s720
func macro_call(def_ref, arg_list, marcro_name pointer) {
    fmt.Println("macro_call not implemented")
}
