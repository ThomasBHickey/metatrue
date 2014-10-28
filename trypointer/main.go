package main

import (
    "fmt"
)

type Node interface {
    setLink(int)
}

type test_tok struct {
    link int
}

// func (node test_tok) setLink(v int) {
//     node.link = v
// }

func (nodep test_tok) setLink(v int){
    nodep.link = v
}

func NewNodes() (toks []*Node){
    return toks
}

var mem = []Node{}

//var memps = []*Node{}
var memps = NewNodes()

func main() {
    fmt.Println("starting main()")
    //mem = append(mem, test_tok{link:7})
    //fmt.Println("mem[0]", mem[0])
    //mem[0].setLink(6)
    //fmt.Println("mem[0] after setLink(6)", mem[0])
    
    memps = append(memps, new(Node))
    fmt.Println("memps[0]", memps[0].(&test_tok))
}
    