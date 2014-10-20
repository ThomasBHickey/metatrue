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

// s782
var (
    job_name string
    log_opened bool
    log_name string
    )
    
// s784
func pack_job_name_sn(sn pointer){
    pack_job_name(pos_to_string[sn])
}

func pack_job_name(s string) {
    cur_area = ""
    cur_ext = s
    cur_name = job_name
    pack_file_name(cur_name, cur_area, cur_ext)
}
    
// s788
func open_log_file(){
    print_err("open_log_file not implemented yet")
    var (
        old_setting int
        k, l, m int
)
    const months = "JANFEBMARAPRMAYJUNJULAUGSEPOCTNOVDEC"
    
    old_setting = selector
    if job_name=="" {job_name = "mfput"}
    pack_job_name(".log")
    for !a_open_out(log_file) {
    }
    
}