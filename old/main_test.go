package metatrue

import "testing"

func TestMain(t *testing.T) {
	err := Main()
	if err != nil {
		t.Errorf("main returned error ", err)
	}
}
