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

// s 24
type eight_bits byte
// don't think we need alpha_file and byte_file

// s 25
// try to do without name_of_file and name_length

// s 26, 27 opening/closing files
// just use os.Open

// s29
var (
	buffer [buf_size + 1] rune
	first,
	last int
	max_buf_stack = 0
)

// s30
func input_ln(f io.Reader, bypass_eoln bool) bool {
    r := bufio.NewReader(f)
    line, err := r.ReadString('\n') //r.ReadBytes('\n')
    if err!=nil && err!=io.EOF{
        return false
    }
    if err==io.EOF && len(line)==0 {
        return false
    }
    //line = bytes.Trim(line, " \t\n")
    runes := []rune(line)
    if len(runes)+last >= max_buf_stack {
        // s 34
        overflow("buffer_size", buf_size)  // this won't return
    }
    last = first
    for _, r := range runes {
        buffer[last] = r
        last++
    }
	return true
}

// s31
var (
	term_in  = os.Stdin
	term_out = os.Stdout
)
//s 32
func t_open_in(){
}

func t_open_out(){
}

// s33
func update_terminal()  {}
func clear_terminal()   {}
func wake_up_terminal() {}

// s36
func init_terminal() error {
	fmt.Println("init_terminal")
	t_open_in()
	fmt.Println("len of args:", len(os.Args), "args:", os.Args[1:])
	if len(os.Args) > 1 {
		rs := ([]rune)(strings.Join(os.Args[1:], " "))
		fmt.Println("bs:", rs)
		for pos, r := range rs {
			buffer[pos] = r
		}
		first = 0
		last = len(rs)
		fmt.Println("init_terminal returning OK")
		return nil
	}
	for {
		wake_up_terminal()
		fmt.Print("**")
		update_terminal()

		fmt.Fprint(term_out, "! End of file on the terminal... why?")
		return errors.New("EOF on terminal")
	}
}
