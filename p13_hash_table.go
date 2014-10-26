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

type eqtb_cell struct{ eq_type, equiv halfword }

var eqtb []eqtb_cell

// s210
func primitive(rs string, c, o halfword) {
	cur_sym = halfword(make_string(rs))
	for int(cur_sym) >= len(eqtb) {
		eqtb = append(eqtb, eqtb_cell{eq_type: 0, equiv: 0})
	}
	eqtb[cur_sym] = eqtb_cell{eq_type: c, equiv: o}
}

// s211
func putPrimitivesIn_s211() {
	primitive("..", path_join, 0)
	primitive("[", left_bracket, 0)
	//eqtb[frozen_left_bracket] = eqtb[cur_sym]
	primitive("]", right_bracket, 0)
	primitive("}", right_brace, 0)
	primitive("{", left_brace, 0)
	primitive(":", colon, 0)
	//eqtb[frozen_colon] = eqtb[cur_sym]
	primitive("::", double_colon, 0)
	primitive("||:", bchar_label, 0)
	primitive(":=", assignment, 0)
	primitive(",", comma, 0)
	primitive(";", semicolon, 0)
	//eqtb[frozen_semicolon] = eqtb[cur_sym]
	primitive("\\", relax, 0)

	primitive("addto", add_to_command, 0)
	primitive("at", at_token, 0)
	primitive("atleast", at_least, 0)
	primitive("begingroup", begin_group, 0)
	bg_loc = cur_sym
	primitive("controls", controls, 0)
	primitive("cull", cull_command, 0)
	primitive("curl", curl_command, 0)
	primitive("delimiters", delimiters, 0)
	primitive("display", display_command, 0)
	primitive("endgroup", end_group, 0)
	//eqtb[frozen_end_group] = eqtb[cur_sym]
	eg_loc = cur_sym
	primitive("everyjob", every_job_command, 0)
	primitive("exitif", exit_test, 0)
	primitive("expandafter", expand_after, 0)
	primitive("from", from_token, 0)
	primitive("inwindow", in_window, 0)
	primitive("interim", interim_command, 0)
	primitive("let", let_command, 0)
	primitive("newinternal", new_internal, 0)
	primitive("of", of_token, 0)
	primitive("openwindow", open_window, 0)
	primitive("randomseed", random_seed, 0)
	primitive("save", save_command, 0)
	primitive("scantokens", scan_tokens, 0)
	primitive("shipout", ship_out_command, 0)
	primitive("skipto", skip_to, 0)
	primitive("step", step_token, 0)
	primitive("str", str_op, 0)
	primitive("tension", tension, 0)
	primitive("to", to_token, 0)
	primitive("until", until_token, 0)
}
