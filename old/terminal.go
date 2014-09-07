package metatrue

import (
	"fmt"
	"os"
)

func t_open_out() {}

func wterm_ln(ss ...string) {
	fmt.Fprintln(os.Stdout, ss)
}
