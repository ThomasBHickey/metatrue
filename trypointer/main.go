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

func (node *test_tok) setLink(v int) {
     node.link = v
 }

func NewNodes() (toks []*Node){
    return toks
}

var mem = []Node{}

func setLink(p, v int){
    mem[p].setLink(v)
}

func main() {
    fmt.Println("starting main()")
    mem = append(mem, &test_tok{link:7})
    //mem = append(mem, test_tok{link:8})
    fmt.Println("mem[0]", mem[0])
    
    node := test_tok{link:9}
    mem[0] = &node
    //mem[0] = &test_tok{link:8} 
    fmt.Println("mem[0]", mem[0])
    
    setLink(0, 6)
    fmt.Println("mem[0] after setLink(6)", mem[0])
    
}
    