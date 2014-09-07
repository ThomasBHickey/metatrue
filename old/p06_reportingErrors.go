package metatrue

// s71, 72
const (
	spotless = iota
	warning_issued
	error_message_issued
	fatal_error_stop
)

var deletions_allowed = true
var history int
var error_count int = 0
