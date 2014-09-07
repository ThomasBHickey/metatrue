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
	"fmt"
)

// s54, 55
const (
	no_print = iota
	term_only
	log_only
	term_and_log
	pseudo
	new_string
	max_selector
)

var term_offset int = 0
var file_offset int = 0

func initializeOutputRoutines() {
	fmt.Println("initializeOutputRoutines")
}
