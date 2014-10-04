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

// s159
type Node interface {
	Type() quarterword
}

var (
	mem     = []Node{}
	free_mem = []pointer{}
	//hi_mem     []two_halves
	//lo_mem_max pointer
	//hi_mem_max pointer // both grow upwards
)

// s160
var var_used, dyn_used integer

// s161
// we hope Go inlines this sort of thing!
func link(p pointer) pointer {
	return pointer(mem[p].(two_halves).rh)
}

func info(mw two_halves) halfword {
	return mw.lh
}

var avail pointer
var mem_end pointer

// s 163
// put node in mem
func get_avail(node Node) pointer {
    var pos pointer
    if len(free_mem)>0 {
        pos = free_mem[len(free_mem)-1]
        free_mem = free_mem[:len(free_mem)-1]
        mem[pos] = node
    }else{
        pos = pointer(len(mem))
        mem = append(mem, node)
    }
    if pos>mem_max {
        runaway()
        overflow("mem memory size", mem_max)
    }
	if stat {
		dyn_used++
	}
	return pos
}

// S172
func free_node(p pointer) {
    free_mem = append(free_mem, p)
}
