package metatrue

import (
	"errors"
)

const (
	srceen_depth    = 1024 // pixels in each column of screen display
	error_line      = 72
	half_error_line = 42
	max_print_line  = 79
	stack_size      = 30  // max # of simultaneous input sources
	screen_width    = 768 // pixels in each row of screen display
	max_in_open     = 6   // max input files
	bistack_size    = 785
)

func check_constants() error {
	bad = 0
	if (half_error_line < 30) || (half_error_line > error_line-15) {
		bad = 1
	}
	if max_print_line < 60 {
		bad = 2
	}
	if bad == 0 {
		return nil
	}
	return errors.New("Check bad")
}
