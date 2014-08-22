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

package metatrue

import (
	"errors"
	"fmt"
	"os"
)

// s29
var (
	buffer []byte
	first,
	last int
	max_buf_stack = 0
)

// s31
var (
	term_in  = os.Stdin
	term_out = os.Stdout
)

// s32
func t_open_in() {}

// s33
func update_terminal(){}
func clear_terminal(){}
func wake_up_terminal(){}

// s36
func init_terminal() error {
	fmt.Println("init_terminal")
	t_open_in()
	for {
		wake_up_terminal()
		fmt.Fprint(term_out, "! End of file on the terminal... why?")
		return errors.New("EOF on terminal")
	}
}
