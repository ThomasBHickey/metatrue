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
	"path/filepath"
)

// s767
var cur_name,
	cur_area,
	cur_ext string

// s780
func make_name_string(f *os.File) {
    fmt.Println("in make_name_string", f)
	fileInfo, err := f.Stat()
	if err != nil {
		fatal_error("fileInfo error in make_name_string")
	}
	fullPath, err := filepath.Abs(fileInfo.Name())
	if err != nil {
		fatal_error("filepath failed in make_name_string")
	}
	name_of_file = fullPath
	fmt.Println("name_of_file in make_name_string", name_of_file)
}

// s774
func pack_file_name(name, area, extension string) {
    name_of_file = ""
    if len(area)>0 {name_of_file = area+"/"}
    name_of_file = name_of_file+name
    if len(extension)>0 {name_of_file = name_of_file+"."+extension}
}

// s782
var (
	job_name   = ""
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
	fmt.Println("pack_job_name", cur_name, cur_ext)
	pack_file_name(cur_name, cur_area, cur_ext)
}

// s788
func open_log_file() {
	fmt.Println("in open_log_file()")
	var (
		old_setting int
		k, l integer //, m     int
	)

	old_setting = selector
	if job_name == "" {
		job_name = "mfput"
	}
	pack_job_name("log")
	for {
	    fmt.Println("in open loop")
		fp := open_out()
		if fp != nil {
			log_file = fp
			break
		}else{
		    fatal_error("open_log_file failed")
		}
		fatal_error("unable to open " + name_of_file)
	}
	fmt.Println("opened up: "+name_of_file)
	//make_name_string(log_file)  // happens in open_out
	log_name = name_of_file
	selector = log_only
	log_opened = true
	print_banner_line()
	fmt.Println("called print_banner_line")
	fmt.Println("input_ptr", input_ptr)
	fmt.Println("len of input_stack", len(input_stack))
	return
	input_stack[input_ptr] = cur_input
	print_nl("**")
	l = input_stack[0].limit - 1
	for k = 1; k <= l; k++ {
		print_char(buffer[k])
	}
	print_ln()
	selector = old_setting + 2 // log_only or term_and_log
}

// s790
func print_banner_line() {
	const months = "JANFEBMARAPRMAYJUNJULAUGSEPOCTNOVDEC"
	fmt.Println("print_banner_line()", banner)
	wlog(banner)
	fmt.Println("passed wlog in print_banner_line")
	slow_print(base_ident)
	fmt.Println("passed slow_print in print_banner_line")
	print("  ")
	fmt.Println("passed double blank in print_banner_line")
	fmt.Println("value of day", day)
	fmt.Println("len(internal)", len(internal))
	fmt.Println("internal[day]", internal[day])
	print_int(round_unscaled(internal[day]))
	fmt.Println("passed print_int in print_banner_line")
	print_char(' ')
	m := round_unscaled(internal[month])
	wlog(months[3*m : 3*m+3])
	print_char(' ')
	print_int(round_unscaled(internal[year]))
	print_char(' ')
	m = round_unscaled(internal[time])
	print_dd(m / 60)
	print_char(':')
	print_dd(m % 60)
	fmt.Println("finishing print_banner_line")
}
