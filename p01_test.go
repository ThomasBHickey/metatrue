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
	"testing"
	//"fmt"
)

func TestX01(t *testing.T) {
	switch true {
	case (half_error_line < 30) || (half_error_line > error_line-15):
		t.Error("bad 1")
	case max_print_line < 60:
		t.Error("bad 2")
	case header_size%4 != 0:
		t.Error("bad 6")
	case (lig_table_size < 255) || (lig_table_size > 32510):
		t.Error("bad 7")
	}
}
