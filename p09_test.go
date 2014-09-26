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

func TestS154(t *testing.T) {
	switch true {
	// mem_max doesn't matter
	case (min_quarterword > 0) || (max_quarterword < 127):
		t.Error("bad 11")
	case (min_halfword > 0) || (max_halfword < 32767):
		t.Error("bad 12")
	case (min_quarterword < min_halfword) || (max_quarterword > max_halfword):
		t.Error("bad 13")
	case (buf_size > max_halfword):
		t.Error("bad 16")
	case (max_quarterword-min_quarterword < 255) || (max_halfword-min_halfword < 65535):
		t.Error("bad 17")
	}
}

func TestS155(t *testing.T){
    if ho(2)!=2 { t.Error("s155 ho failed")}
    if qo(2)!=2 { t.Error("s155 qo failed")}
    if qi(2)!=2 { t.Error("s155 qi failed")}
}
