package main

import (
    "fmt"
)

type (
    small_number byte
    quarterword byte
    pointer uint16
    str_number pointer
    halfword uint16
    scaled int32
    )

type Node interface {
	getType() small_number
	setLink(pointer)
	getLink() pointer
	getName_type() quarterword
}

type num_tok struct {
	value     scaled
	link      pointer
	info      halfword
	name_type quarterword
}

func (node *num_tok) getType() small_number {
	return 123
}

func (node *num_tok) getName_type() quarterword {
    return node.name_type
}

func (node *num_tok) setLink(l pointer){
	node.link = l
}

func (node *num_tok) getLink() pointer{
	return node.link
}

// Now for something a bit different

type InputStateRec interface {
    getIndex() quarterword
    getName() str_number
    setIndex(quarterword)
}

type inStateFileRec struct {
	index quarterword
	start,
	loc,
	limit halfword
	name str_number
}

func (input_state inStateFileRec) getIndex() quarterword {
    return input_state.index
}

func (input_state inStateFileRec) getName() str_number {
    return input_state.name
}

func (input_state *inStateFileRec) setIndex(ndx quarterword) {
    input_state.index = ndx
}

const stack_size=32

var (
	mem     = []Node{}
	input_stack  [stack_size + 1]InputStateRec
	cur_input InputStateRec
	)
func main() {
    fmt.Println("starting main()")
    mem = append(mem, &num_tok{link:7})
    fmt.Println("mem[0]", mem[0])
    mem[0].setLink(6)
    fmt.Println("mem[0] after setLink(6)", mem[0])
    input_stack[0] = &inStateFileRec{index:101, name:15}
    fmt.Println("input_stack[0]", input_stack[0])
    input_stack[0].setIndex(99)
    fmt.Printf("input_stack[0]: %#v\n", input_stack[0])
    test := input_stack[0].(*inStateFileRec)
    test.loc = 7
    fmt.Printf("test %#v\n", test)
    cur_input = input_stack[0]
    //test = cur_input.(*inStateFileRec)
    cur_input.(*inStateFileRec).loc = 8
    fmt.Printf("input_stack[0]: %#v\n", input_stack[0])
    }
    