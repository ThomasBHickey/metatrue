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
	"bufio"
	//"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func write(w *bufio.Writer, msgs ...string) {
	//fmt.Println("in write()", msgs)
	for _, msg := range msgs {
		//fmt.Println("msg in msg loop:", msg)
		_, err := w.WriteString(msg)
		if err != nil {
			jump_out(err)
		}
	}
}

func write_ln(w *bufio.Writer, msgs ...string) {
	for _, msg := range msgs {
		_, err := w.WriteString(msg)
		if err != nil {
			jump_out(err)
		}
	}
	w.WriteRune('\n')
	w.Flush()
}

// s 24
type eight_bits byte

// don't think we need alpha_file and byte_file

// s 25
var name_of_file string

// s 26

func open_out() *bufio.Writer {
	fmt.Println("open_out() ", name_of_file)
	fo, err := os.Create(name_of_file)
	fmt.Println("fo, err", fo, err)
	if err != nil {
		print_err("problem opening in open_out()")
		jump_out(err)
	}
	make_name_string(fo)
	w := bufio.NewWriter(fo)
	return w
}

// s29
var (
	buffer [buf_size + 1]rune
	first,
	last halfword
	max_buf_stack halfword
)

// s30
func input_ln(r *bufio.Reader, bypass_eoln bool) bool {
    fmt.Println("starting input_ln")
	line, err := r.ReadString('\n') //r.ReadBytes('\n')
	line = strings.TrimSpace(line)
	fmt.Printf("input_ln s30 line: %s \"%s\"\n", len(line), line)
	fmt.Println("err", err)
	if err != nil && err != io.EOF {
		return false
	}
	if line=="!q" || line=="!end" ||line=="quit()" {
	    fmt.Println("found end, exiting")
	    jump_out(errors.New("input_ln in p03 exiting on 'end'"))
	}
	if err == io.EOF && len(line) == 0 {
		return false
	}
	//line = bytes.Trim(line, " \t\n")
	runes := []rune(line)
	fmt.Printf("len(runes) %d, last: %d, max_buf_stack: %d\n", len(runes), int(last), max_buf_stack)
	if len(runes)+int(last) >= buf_size {
		// s 34
		fmt.Println("input_ln about to call overflow")
		overflow("buffer_size", buf_size) // this won't return
	}
	last = first
	for _, r := range runes {
		buffer[last] = r
		last++
	}
	max_buf_stack = last
	return true
}

// s31
var (
	term_in  = bufio.NewReader(os.Stdin)
	term_out = bufio.NewWriter(os.Stdout)
)

//s 32
func t_open_in() {
}

func t_open_out() *bufio.Writer {
	return bufio.NewWriter(os.Stdout)
}

// s33
func update_terminal() {
	term_out.Flush()
}
func clear_terminal()   {}
func wake_up_terminal() {}

// s36
func bufferText(s string) {
	s = strings.Trim(s, " \n\t")
	rs := ([]rune)(s)
	for pos, r := range rs {
		buffer[pos] = r
	}
	fmt.Printf("in bufferText %s %#v", s, cur_input)
	if cur_input == nil {
		jump_out(errors.New("nil cur_input in bufferText()"))
	}
	fmt.Printf("in bufferText %s %#v", s, cur_input)
	cur_input.(*inStateFileRec).loc = first

	last = halfword(len(rs))
}

func init_terminal() error {
	fmt.Println("init_terminal")
	t_open_in()
	//fmt.Println("len of args:", len(os.Args), "args:", os.Args[1:])
	if len(os.Args) > 1 {
		text := strings.Join(os.Args[1:], " ")
		bufferText(text)
		if cur_input.(*inStateFileRec).loc < last {
			return nil
		}

	}
	for {
		wake_up_terminal()
		fmt.Print("**")
		update_terminal()
		text, err := term_in.ReadString('\n')
		if err != nil {
			fmt.Fprintln(term_out, "! End of file on the terminal... why?")
			return err
		}
		bufferText(text)
		if cur_input.(*inStateFileRec).loc < last {
			return nil
		}
	}
}
