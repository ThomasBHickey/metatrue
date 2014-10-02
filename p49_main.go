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
// s1203
//var ready_already int

// s1204
func Start() error {
	var err error = nil
	fmt.Println("MT starting")
	history = fatal_error_stop // in case we quit during initialization
	t_open_out()
// 	if ready_already == 314159 {
// 		goto start_of_MT
// 	}
// 	err = check_constants()
// 	if err != nil {
// 		wterm_ln("Ouch--my internal constants have been clobbered!", "---case ", err.Error())
// 		goto final_end
// 	}
	initialize()
	//ready_already = 314159
//start_of_MT:
	fmt.Println("mt start_of_MT")
	if err != nil {
		return err
	}
	initialize_output_routines()
	fmt.Println("start_of_MT")
	err = getFirstLineOfInputAndPrepareToStart()
	if err != nil {
		return err
	}
	fmt.Println("first:", first,", last:", last)
	fmt.Println("first line of input", string(buffer[cur_input.loc:last]))
	history = spotless
	if start_sym > 0 {
		cur_sym = start_sym
		back_input()
	}
	//goto end_of_MT
	main_control()
	err = final_cleanup()
	if err != nil {
		fmt.Println("error calling final_cleanup:", err)
		goto final_end
	}
	//end_of_MT:
	err = close_files_and_terminate()
	if err != nil {
		fmt.Println("error calling close_files_and_terminate:", err)
		goto final_end
	}

final_end:
	fmt.Println("at final_end")
	return err
}

// s1205
func close_files_and_terminate() error {
	fmt.Println("close_files_and_terminate")
	return nil
}

// s 1209
func final_cleanup() error {
	fmt.Println("final_cleanup")
	return nil
}

// Section 1211
func getFirstLineOfInputAndPrepareToStart() error {
	fmt.Println("getFirstLineOfInputAndPrepareToStart()")
	return initializeTheInputRoutines()
}
