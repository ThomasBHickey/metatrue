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
    "time"
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

type command_code int

// s187

const (
	undefined        = 0 //{no type has been declared}
	unknown_tag      = 1 //{this constant is added to certain type codes below}
	vacuous          = 1 //{no expression was present}
	boolean_type     = 2 //{\&{boolean} with a known value}
	unknown_boolean  = boolean_type + unknown_tag
	string_type      = 4 //{\&{string} with a known value}
	unknown_string   = string_type + unknown_tag
	pen_type         = 6 //{\&{pen} with a known value}
	unknown_pen      = pen_type + unknown_tag
	future_pen       = 8 //{subexpression that will become a \&{pen} at a higher level}
	path_type        = 9 //{\&{path} with a known value}
	unknown_path     = path_type + unknown_tag
	picture_type     = 11 //{\&{picture} with a known value}
	unknown_picture  = picture_type + unknown_tag
	transform_type   = 13 //{\&{transform} variable or capsule}
	pair_type        = 14 //{\&{pair} variable or capsule}
	numeric_type     = 15 //{variable that has been declared \&{numeric} but not used}
	known            = 16 //{\&{numeric} with a known value}
	dependent        = 17 //{a linear combination with |fraction| coefficients}
	proto_dependent  = 18 //{a linear combination with |scaled| coefficients}
	independent      = 19 //{\&{numeric} with unknown value}
	token_list       = 20 //{variable name or suffix argument or text argument}
	structured       = 21 //{variable with subscripts and attributes}
	unsuffixed_macro = 22 //{variable defined with \&{vardef} but no \.{\AT!\#}}
	suffixed_macro   = 23 //{variable defined with \&{vardef} and \.{\AT!\#}}
	symbolic         = 24 // added for single word nodes Th
)

var unknown_types = [...]int{unknown_boolean, unknown_string,
	unknown_pen, unknown_picture, unknown_path}

func print_type(t small_number) {
	switch t {
	case vacuous:
		print("vacuous")
	case boolean_type:
		print("boolean")
	case unknown_boolean:
		print("unknown boolean")
	case string_type:
		print("string")
	case unknown_string:
		print("unknown string")
	case pen_type:
		print("pen")
	case unknown_pen:
		print("unknown pen")
	case future_pen:
		print("future pen")
	case path_type:
		print("path")
	case unknown_path:
		print("unknown path")
	case picture_type:
		print("picture")
	case unknown_picture:
		print("unknown picture")
	case transform_type:
		print("transform")
	case pair_type:
		print("pair")
	case known:
		print("known numeric")
	case dependent:
		print("dependent")
	case proto_dependent:
		print("proto-dependent")
	case numeric_type:
		print("numeric")
	case independent:
		print("independent")
	case token_list:
		print("token list")
	case structured:
		print("structured")
	case unsuffixed_macro:
		print("unsuffixed macro")
	case suffixed_macro:
		print("suffixed macro")
	default:
		print("undefined")
	}
}

// s188
const (
	root = iota
	saved_root
	structured_root
	subscr
	attr
	x_part_sector
	y_part_sector
	xx_part_sector
	xy_part_sector
	yx_part_sector
	yy_part_sector
	capsule
	token
)

// s189
const (
	true_code = iota + 30
	false_code
	null_picture_code
	null_pen_code
	job_name_op
	read_string_op
	pen_circle
	normal_deviate
	odd_op
	known_op
	unknown_op
	not_op
	decimal
	reverse
	make_path_op
	make_pen_op
	total_weight_op
	oct_op
	hex_op
	ASCII_op
	char_op
	length_op
	turning_op
	x_part
	y_part
	xx_part
	xy_part
	yx_part
	yy_part
	sqrt_op
	m_exp_op
	m_log_op
	sin_d_op
	cos_d_op
	floor_op
	uniform_deviate
	char_exists_op
	angle_op
	cycle_op
	plus
	minus
	times
	over
	pythag_add
	pythag_sub
	or_op
	and_op
	less_than
	less_or_equal
	greater_than
	greater_or_equal
	equal_to
	unequal_to
	concatenate
	rotated_by
	slanted_by
	scaled_by
	shifted_by
	transformed_by
	x_scaled
	y_scaled
	z_scaled
	intersect
	double_dot
	substring_of
	//min_of
	subpath_of
	direction_time_of
	point_of
	precontrol_of
	postcontrol_of
	pen_offset_of
	min_of = substring_of
)

func print_op(c quarterword) {
	if c <= numeric_type {
		print_type(small_number(c))
	} else {
		switch c {
		case true_code:
			print("true")
		case false_code:
			print("false")
		case null_picture_code:
			print("nullpicture")
		case null_pen_code:
			print("nullpen")
		case job_name_op:
			print("jobname")
		case read_string_op:
			print("readstring")
		case pen_circle:
			print("pencircle")
		case normal_deviate:
			print("normaldeviate")
		case odd_op:
			print("odd")
		case known_op:
			print("known")
		case unknown_op:
			print("unknown")
		case not_op:
			print("not")
		case decimal:
			print("decimal")
		case reverse:
			print("reverse")
		case make_path_op:
			print("makepath")
		case make_pen_op:
			print("makepen")
		case total_weight_op:
			print("totalweight")
		case oct_op:
			print("oct")
		case hex_op:
			print("hex")
		case ASCII_op:
			print("ASCII")
		case char_op:
			print("char")
		case length_op:
			print("length")
		case turning_op:
			print("turningnumber")
		case x_part:
			print("xpart")
		case y_part:
			print("ypart")
		case xx_part:
			print("xxpart")
		case xy_part:
			print("xypart")
		case yx_part:
			print("yxpart")
		case yy_part:
			print("yypart")
		case sqrt_op:
			print("sqrt")
		case m_exp_op:
			print("mexp")
		case m_log_op:
			print("mlog")
		case sin_d_op:
			print("sind")
		case cos_d_op:
			print("cosd")
		case floor_op:
			print("floor")
		case uniform_deviate:
			print("uniformdeviate")
		case char_exists_op:
			print("charexists")
		case angle_op:
			print("angle")
		case cycle_op:
			print("cycle")
		case plus:
			print_char('+')
		case minus:
			print_char('-')
		case times:
			print_char('*')
		case over:
			print_char('/')
		case pythag_add:
			print("++")
		case pythag_sub:
			print("+-+")
		case or_op:
			print("or")
		case and_op:
			print("and")
		case less_than:
			print_char('<')
		case less_or_equal:
			print("<=")
		case greater_than:
			print_char('>')
		case greater_or_equal:
			print(">=")
		case equal_to:
			print_char('=')
		case unequal_to:
			print("<>")
		case concatenate:
			print("&")
		case rotated_by:
			print("rotated")
		case slanted_by:
			print("slanted")
		case scaled_by:
			print("scaled")
		case shifted_by:
			print("shifted")
		case transformed_by:
			print("transformed")
		case x_scaled:
			print("xscaled")
		case y_scaled:
			print("yscaled")
		case z_scaled:
			print("zscaled")
		case intersect:
			print("intersectiontimes")
		case substring_of:
			print("substring")
		case subpath_of:
			print("subpath")
		case direction_time_of:
			print("directiontime")
		case point_of:
			print("point")
		case precontrol_of:
			print("precontrol")
		case postcontrol_of:
			print("postcontrol")
		case pen_offset_of:
			print("penoffset")
		default:
			print("..")
		}
	}
}

// s190
const (
	tracing_titles = iota + 1
	tracing_equations
	tracing_capsules
	tracing_choices
	tracing_specs
	tracing_pens
	tracing_commands
	tracing_restores
	tracing_macros
	tracing_edges
	tracing_output
	tracing_stats
	tracing_online
	year
	month
	day
	Time
	char_code
	char_ext
	char_wd
	char_ht
	char_dp
	char_ic
	char_dx
	char_dy
	design_size
	hppp
	vppp
	x_offset
	y_offset
	pausing
	showstopping
	fontmaking
	proofing
	smoothing
	autorounding
	granularity
	fillin
	turning_check
	warning_check
	boundary_char
	max_given_internal = boundary_char
)

var (
	internal []scaled
	int_name []str_number
	internals_set_up bool
)

// s191
func set_up_internals() {
    fmt.Println("in set_up_internals()", internals_set_up)
    if internals_set_up { return }
	for k := 0; k <= max_given_internal; k++ {
		internal = append(internal, 0)
	}
	internals_set_up = true
}

// s192
func putPrimitivesIn_s192() {
	primitive("tracingtitles", internal_quantity, tracing_titles)
	primitive("tracingequations", internal_quantity, tracing_equations)
	primitive("tracingcapsules", internal_quantity, tracing_capsules)
	primitive("tracingchoices", internal_quantity, tracing_choices)
	primitive("tracingspecs", internal_quantity, tracing_specs)
	primitive("tracingpens", internal_quantity, tracing_pens)
	primitive("tracingcommands", internal_quantity, tracing_commands)
	primitive("tracingrestores", internal_quantity, tracing_restores)
	primitive("tracingmacros", internal_quantity, tracing_macros)
	primitive("tracingedges", internal_quantity, tracing_edges)
	primitive("tracingoutput", internal_quantity, tracing_output)
	primitive("tracingstats", internal_quantity, tracing_stats)
	primitive("tracingonline", internal_quantity, tracing_online)
	primitive("year", internal_quantity, year)
	primitive("month", internal_quantity, month)
	primitive("day", internal_quantity, day)
	primitive("time", internal_quantity, Time)
	primitive("charcode", internal_quantity, char_code)
	primitive("charext", internal_quantity, char_ext)
	primitive("charwd", internal_quantity, char_wd)
	primitive("charht", internal_quantity, char_ht)
	primitive("chardp", internal_quantity, char_dp)
	primitive("charic", internal_quantity, char_ic)
	primitive("chardx", internal_quantity, char_dx)
	primitive("chardy", internal_quantity, char_dy)
	primitive("designsize", internal_quantity, design_size)
	primitive("hppp", internal_quantity, hppp)
	primitive("vppp", internal_quantity, vppp)
	primitive("xoffset", internal_quantity, x_offset)
	primitive("yoffset", internal_quantity, y_offset)
	primitive("pausing", internal_quantity, pausing)
	primitive("showstopping", internal_quantity, showstopping)
	primitive("fontmaking", internal_quantity, fontmaking)
	primitive("proofing", internal_quantity, proofing)
	primitive("smoothing", internal_quantity, smoothing)
	primitive("autorounding", internal_quantity, autorounding)
	primitive("granularity", internal_quantity, granularity)
	primitive("fillin", internal_quantity, fillin)
	primitive("turningcheck", internal_quantity, turning_check)
	primitive("warningcheck", internal_quantity, warning_check)
	primitive("boundarychar", internal_quantity, boundary_char)
}

// s193
func initialize_table_entries_A() {
    fmt.Println("initialize_table_entries_A")
	// int_name has been pre-extended to accommodate these entries
	int_name[tracing_titles] = make_string("tracingtitles")
	int_name[tracing_equations] = make_string("tracingequations")
	int_name[tracing_capsules] = make_string("tracingcapsules")
	int_name[tracing_choices] = make_string("tracingchoices")
	int_name[tracing_specs] = make_string("tracingspecs")
	int_name[tracing_pens] = make_string("tracingpens")
	int_name[tracing_commands] = make_string("tracingcommands")
	int_name[tracing_restores] = make_string("tracingrestores")
	int_name[tracing_macros] = make_string("tracingmacros")
	int_name[tracing_edges] = make_string("tracingedges")
	int_name[tracing_output] = make_string("tracingoutput")
	int_name[tracing_stats] = make_string("tracingstats")
	int_name[tracing_online] = make_string("tracingonline")
	int_name[year] = make_string("year")
	int_name[month] = make_string("month")
	int_name[day] = make_string("day")
	int_name[Time] = make_string("time")
	int_name[char_code] = make_string("charcode")
	int_name[char_ext] = make_string("charext")
	int_name[char_wd] = make_string("charwd")
	int_name[char_ht] = make_string("charht")
	int_name[char_dp] = make_string("chardp")
	int_name[char_ic] = make_string("charic")
	int_name[char_dx] = make_string("chardx")
	int_name[char_dy] = make_string("chardy")
	int_name[design_size] = make_string("designsize")
	int_name[hppp] = make_string("hppp")
	int_name[vppp] = make_string("vppp")
	int_name[x_offset] = make_string("xoffset")
	int_name[y_offset] = make_string("yoffset")
	int_name[pausing] = make_string("pausing")
	int_name[showstopping] = make_string("showstopping")
	int_name[fontmaking] = make_string("fontmaking")
	int_name[proofing] = make_string("proofing")
	int_name[smoothing] = make_string("smoothing")
	int_name[autorounding] = make_string("autorounding")
	int_name[granularity] = make_string("granularity")
	int_name[fillin] = make_string("fillin")
	int_name[turning_check] = make_string("turningcheck")
	int_name[warning_check] = make_string("warningcheck")
	int_name[boundary_char] = make_string("boundarychar")
}

// s194
func fix_date_and_time() {
    now := time.Now()
    cyear, cmonth, cday := now.Date()
    chour := now.Hour()
    cminute := now.Minute()
    //csecond := now.Second()
	internal[Time] = scaled(chour*60 + cminute) * unity // 12 * 60 * unity
	internal[day] = scaled(cday) * unity // 4 * unity
	internal[month] = scaled(cmonth) * unity //7 * unity
	internal[year] = scaled(cyear) * unity //1776 * unity
	fmt.Println("internal[day]", internal[day])
}

// s195
func begin_diagnostic() {
	old_setting = selector
	if (internal[tracing_online] == 0) || (selector == term_and_log) {
		selector--
		if history == spotless {
			history = warning_issued
		}
	}
}

func end_diagnostic(blank_line bool) {
	print_nl("")
	if blank_line {
		print_ln()
	}
	selector = old_setting
}

// s196
var old_setting int

// s197
func print_diagnostic(s, t str_number, nuline bool) {
	begin_diagnostic()
	if nuline {
		print_nl_sn(s)
	} else {
		print_sn(s)
	}
	print(" at line ")
	print_int(line)
	print_sn(t)
	print_char(':')
}

// s198
const (
	digit_class = iota
	period_class
	space_class
	percent_class
	string_class
	right_paren_class = 8
	//isolated_classes = {5,6,7,8}
	letter_class        = 9
	left_bracket_class  = 17
	right_bracket_class = 18
	invalid_class       = 20
	max_class           = 20
)

var isolated_classes = [...]int{5, 6, 7, 8}
var char_class [256]byte

// s199
// Eventually we'll allow any Unicode chars, at least in variable
// names, etc.

func setup_char_class() {
	fmt.Println("setup_char_class()")
	for k := 0; k < 256; k++ {
		char_class[k] = invalid_class
	}
	for k := '0'; k <= '9'; k++ {
		char_class[k] = digit_class
	}
	char_class['.'] = period_class
	char_class[' '] = space_class
	char_class['%'] = percent_class
	char_class['"'] = string_class
	char_class['_'] = letter_class
	char_class['<'] = 10
	char_class['='] = 10
	char_class['>'] = 10
	char_class[':'] = 10
	char_class['`'] = 11
	char_class['\''] = 11
	char_class['+'] = 12
	char_class['-'] = 12
	char_class['/'] = 13
	char_class['*'] = 13
	char_class['\\'] = 13
	char_class['!'] = 14
	char_class['?'] = 14
	char_class['#'] = 15
	char_class['&'] = 15
	char_class['@'] = 15
	char_class['$'] = 15
	char_class['^'] = 16
	char_class['~'] = 16
	char_class['['] = left_bracket_class
	char_class[']'] = right_bracket_class
	char_class['{'] = 19
	char_class['}'] = 19
// 	for k := 0; k <= ' '-1; k++ {
// 		char_class[k] = invalid_class
// 	}
// 	for k := 127; k <= 255; k++ {
// 		char_class[k] = invalid_class
// 	}
	fmt.Println("rune and char_class of backslash: ", '\\', char_class['\\'])
}
