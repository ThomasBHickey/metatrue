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


// The dynamic memory in MetaTrue has been reorganized from MetaFont
//  The two upper and lower regions are now maintained in separate 
// Go slices.

// s158
type pointer halfword
const null = mem_min
type himemNode struct{
}

type (
    node2 struct {w1, w2 integer}
    node3 struct {w1, w2, w3 integer}
    allnode interface{}
    )



// s159
var (
    lomem = make([]interface{}, 1000)
    himem []two_halves
    lo_mem_max pointer
    hi_mem_max pointer  // himem grows upwards too
    )
    
func test() {
    n2p := *&node2{1,2}
    //n3p := &node3{1,2,3}
    lomem.append(lomem, interface{}(n2p))
}

// s160
var var_used, dyn_used integer

// s161
// we hope Go inlines this sort of thing!
func link(mw two_halves) halfword {
    return mw.rh
}

func info(mw two_halves) halfword{
    return mw.lh
}
