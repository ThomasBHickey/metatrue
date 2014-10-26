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

// s683
const (
    start_def = 1
    var_def = 2
    end_def = 0
    start_forever = 1
    end_for = 0
    )

func putPrimitivesIn_s683() {
    primitive("def", macro_def, start_def)
    primitive("vardef", macro_def, var_def)
    primitive("primarydef", macro_def, secondary_primary_macro)
    primitive("secondarydef", macro_def, tertiary_secondary_macro)
    primitive("tertiarydef", macro_def, expression_tertiary_macro)
    primitive("enddef", macro_def, end_def)
    //eqtb[frozen_end_def] = eqtb[cur_sym]
    primitive("for", iteration, expr_base)
    primitive("forsuffixes", iteration, suffix_base)
    primitive("forever", iteration, start_forever)
    primitive("endfor", iteration, end_for)
}

// s699
var bg_loc, eg_loc halfword
    