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

// s186, mostly pulled in directly from the MF source
const (
	if_test                   = 1  //{conditional text (\&//{if})}
	fi_or_else                = 2  //{delimiters for conditionals (\&//{elseif}, \&//{else}, \&//{fi})}
	input                     = 3  //{input a source file (\&//{input}, \&//{endinput})}
	iteration                 = 4  //{iterate (\&//{for}, \&//{forsuffixes}, \&//{forever}, \&//{endfor})}
	repeat_loop               = 5  //{special command substituted for \&//{endfor}}
	exit_test                 = 6  //{premature exit from a loop (\&//{exitif})}
	relax                     = 7  //{do nothing (\.//{\char`\\})}
	scan_tokens               = 8  //{put a string into the input buffer}
	expand_after              = 9  //{look ahead one token}
	defined_macro             = 10 //{a macro defined by the user}
	min_command               = defined_macro + 1
	display_command           = 11 //{online graphic output (\&//{display})}
	save_command              = 12 //{save a list of tokens (\&//{save})}
	interim_command           = 13 //{save an internal quantity (\&//{interim})}
	let_command               = 14 //{redefine a symbolic token (\&//{let})}
	new_internal              = 15 //{define a new internal quantity (\&//{newinternal})}
	macro_def                 = 16 //{define a macro (\&//{def}, \&//{vardef}, etc.)}
	ship_out_command          = 17 //{output a character (\&//{shipout})}
	add_to_command            = 18 //{add to edges (\&//{addto})}
	cull_command              = 19 //{cull and normalize edges (\&//{cull})}
	tfm_command               = 20 //{command for font metric info (\&//{ligtable}, etc.)}
	protection_command        = 21 //{set protection flag (\&//{outer}, \&//{inner})}
	show_command              = 22 //{diagnostic output (\&//{show}, \&//{showvariable}, etc.)}
	mode_command              = 23 //{set interaction level (\&//{batchmode}, etc.)}
	random_seed               = 24 //{initialize random number generator (\&//{randomseed})}
	message_command           = 25 //{communicate to user (\&//{message}, \&//{errmessage})}
	every_job_command         = 26 //{designate a starting token (\&//{everyjob})}
	delimiters                = 27 //{define a pair of delimiters (\&//{delimiters})}
	open_window               = 28 //{define a window on the screen (\&//{openwindow})}
	special_command           = 29 //{output special info (\&//{special}, \&//{numspecial})}
	type_name                 = 30 //{declare a type (\&//{numeric}, \&//{pair}, etc.)}
	max_statement_command     = type_name
	min_primary_command       = type_name
	left_delimiter            = 31 //{the left delimiter of a matching pair}
	begin_group               = 32 //{beginning of a group (\&//{begingroup})}
	nullary                   = 33 //{an operator without arguments (e.g., \&//{normaldeviate})}
	unary                     = 34 //{an operator with one argument (e.g., \&//{sqrt})}
	str_op                    = 35 //{convert a suffix to a string (\&//{str})}
	cycle                     = 36 //{close a cyclic path (\&//{cycle})}
	primary_binary            = 37 //{binary operation taking `\&//{of}' (e.g., \&//{point})}
	capsule_token             = 38 //{a value that has been put into a token list}
	string_token              = 39 //{a string constant (e.g., |"hello"|)}
	internal_quantity         = 40 //{internal numeric parameter (e.g., \&//{pausing})}
	min_suffix_token          = internal_quantity
	tag_token                 = 41 //{a symbolic token without a primitive meaning}
	numeric_token             = 42 //{a numeric constant (e.g., \.//{3.14159})}
	max_suffix_token          = numeric_token
	plus_or_minus             = 43            //{either `\.+' or `\.-'}
	max_primary_command       = plus_or_minus //{should also be |numeric_token+1|}
	min_tertiary_command      = plus_or_minus
	tertiary_secondary_macro  = 44 //{a macro defined by \&//{secondarydef}}
	tertiary_binary           = 45 //{an operator at the tertiary level (e.g., `\.//{++}')}
	max_tertiary_command      = tertiary_binary
	left_brace                = 46 //{the operator `\.//{\char`\//{}'}
	min_expression_command    = left_brace
	path_join                 = 47 //{the operator `\.//{..}'}
	ampersand                 = 48 //{the operator `\.\&'}
	expression_tertiary_macro = 49 //{a macro defined by \&//{tertiarydef}}
	expression_binary         = 50 //{an operator at the expression level (e.g., `\.<')}
	equals                    = 51 //{the operator `\.='}
	max_expression_command    = equals
	and_command               = 52 //{the operator `\&//{and}'}
	min_secondary_command     = and_command
	secondary_primary_macro   = 53 //{a macro defined by \&//{primarydef}}
	slash                     = 54 //{the operator `\./'}
	secondary_binary          = 55 //{an operator at the binary level (e.g., \&//{shifted})}
	max_secondary_command     = secondary_binary
	param_type                = 56 //{type of parameter (\&//{primary}, \&//{expr}, \&//{suffix}, etc.)}
	controls                  = 57 //{specify control points explicitly (\&//{controls})}
	tension                   = 58 //{specify tension between knots (\&//{tension})}
	at_least                  = 59 //{bounded tension value (\&//{atleast})}
	curl_command              = 60 //{specify curl at an end knot (\&//{curl})}
	macro_special             = 61 //{special macro operators (\&//{quote}, \.//{\#\AT!}, etc.)}
	right_delimiter           = 62 //{the right delimiter of a matching pair}
	left_bracket              = 63 //{the operator `\.['}
	right_bracket             = 64 //{the operator `\.]'}
	right_brace               = 65 //{the operator `\.//{\char`\}}'}
	with_option               = 66 //{option for filling (\&//{withpen}, \&//{withweight})}
	cull_op                   = 67 //{the operator `\&//{keeping}' or `\&//{dropping}'}
	thing_to_add              = 68
	//{variant of \&//{addto} (\&//{contour}, \&//{doublepath}, \&//{also})}
	of_token       = 69 //{the operator `\&//{of}'}
	from_token     = 70 //{the operator `\&//{from}'}
	to_token       = 71 //{the operator `\&//{to}'}
	at_token       = 72 //{the operator `\&//{at}'}
	in_window      = 73 //{the operator `\&//{inwindow}'}
	step_token     = 74 //{the operator `\&//{step}'}
	until_token    = 75 //{the operator `\&//{until}'}
	lig_kern_token = 76
	//{the operators `\&//{kern}' and `\.//{=:}' and `\.//{=:\char'174}', etc.}
	assignment   = 77 //{the operator `\.//{:=}'}
	skip_to      = 78 //{the operation `\&//{skipto}'}
	bchar_label  = 79 //{the operator `\.//{\char'174\char'174:}'}
	double_colon = 80 //{the operator `\.//{::}'}
	colon        = 81 //{the operator `\.:'}

	comma = 82 //{the operator `\.,', must be |colon+1|}
	//end_of_statement==cur_cmd>comma
	semicolon        = 83 ////{the operator `\.;', must be |comma+1|}
	end_group        = 84 //{end a group (\&//{endgroup}), must be |semicolon+1|}
	stop             = 85 //{end a job (\&//{end}, \&//{dump}), must be |end_group+1|}
	max_command_code = stop
	outer_tag        = max_command_code + 1 //{protection code added to command code}
)

func end_of_statement() bool {
	return cur_cmd > comma
}
