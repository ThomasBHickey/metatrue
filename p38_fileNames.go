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
    "os"
    "path/filepath"
)

// s767
var cur_name,
	cur_area,
	cur_ext string
	
// s780
func make_name_string(f *os.File) {
    fileInfo, err := f.Stat()
     if err != nil {
         fatal_error("fileInfo error in make_name_string")
     }
     fullPath, err :=  filepath.Abs(fileInfo.Name())
     if err != nil {
         fatal_error("filepath failed in make_name_string")
     }
     name_of_file = fullPath
}

// s774
func pack_file_name(name, area, extension string) {
    name_of_file = area+"/"+name+"."+extension
    }
    

// s782
var (
	job_name   string
	log_opened bool
	log_name   string
)

// s784
func pack_job_name_sn(sn pointer) {
	pack_job_name(pos_to_string[sn])
}

func pack_job_name(s string) {
	cur_area = ""
	cur_ext = s
	cur_name = job_name
	pack_file_name(cur_name, cur_area, cur_ext)
}

// s788
func open_log_file() {
	print_err("open_log_file not implemented yet")
	var (
		old_setting int
		//k, l, m     int
	)

	old_setting = selector
	if job_name == "" {
		job_name = "mfput"
	}
	pack_job_name(".log")
	for {
	    fp  := open_out()
	    if fp!=nil {
	        log_file = fp
	        break}
	    fatal_error("unable to open "+name_of_file)
	}
	log_name = make_name_string(log_file)
	selector = log_only
	log_opened = true
	print_banner_line()
	input_stack[input_ptr] = cur_input
	print_nl("**")
	l := input_stack[0].limit_field-1
	print(buffer)
	print_ln()
	selector = old_setting + 2 // log_only or term_and_log
}

// s790
func print_banner_line(){
	const months = "JANFEBMARAPRMAYJUNJULAUGSEPOCTNOVDEC"
    wlog(banner)
    slow_print(base_ident)
    print("  ")
    print_int(round_unscaled(internal[day]))
    print_char(' ')
    m := round_unscaled(internal[month])
    wlog(months[3*m:3*m+3])
    print_char(' ')
    print_int(round_unscaled(internal[year]))
    print_char(' ')
    m = round_unscaled(internal[time])
    print_dd(m/60)
    print_char(':')
    print_dd(m%60)
}
