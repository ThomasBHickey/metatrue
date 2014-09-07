package main

import (
	"fmt"
	"io/ioutil"
)

var header = `//	Copyright 2014 Thomas B. Hickey (thomasbhickey@gmail.com)
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
    "testing"
    //"fmt"
)

func TestX(t *testing.T) {
}
`
var fileNames = []string{
	"p01_Inroduction",
	"p02_character_set",
	"p03_input_and_output",
	"p04_string_handling",
	"p05_printing",
	"p06_reporting_errors",
	"p07_scaled_arithmetic",
	"p08_algebric_transcendental_functions",
	"p09_packed_data",
	"p10_dynamic_allocation",
	"p11_memory_layout",
	"p12_command_codes",
	"p13_hash_table",
	"p14_token_lists",
	"p15_variable_data_structures",
	"p16_saving_restoring_equivalents",
	"p17_path_data_structures",
	"p18_choosing_control_points",
	"p19_generating_discrete_moves",
	"p20_edge_structures",
	"p21_octant_subdivision",
	"p22_countour_filling",
	"p23_polygonal_pens",
	"p24_envelope_filling",
	"p25_elliptical_pens",
	"p26_direction_intersection",
	"p27_online_graphic",
	"p28_dynamic_linear_equations",
	"p29_dynamic_nonlinear_equations",
	"p30_syntatic_routines_intro",
	"p31_input_stacks_states",
	"p32_maintaining_input_stacks",
	"p33_getting_next_token",
	"p34_scanning_macro_definitions",
	"p35_exapanding_next_token",
	"p36_conditional_processing",
	"p37_iterations",
	"p38_fileNames",
	"p39_parsing_intro",
	"p40_parsing_primary",
	"p41_parsing_higher",
	"p42_doing_operations",
	"p43_statements_control",
	"p44_commands",
	"p45_font_metric_data",
	"p46_generic_font_format",
	"p47_shipping_characters",
	"p48_dumping_tables",
	"p49_main",
	"p50_debugging",
}

func main() {
	for _, name := range fileNames {
		fmt.Println(name[:4] + "test.go")
		ioutil.WriteFile(name[:4]+"test.go", []byte(header), 0600)
	}
}
