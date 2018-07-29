package bot

import "testing"

func TestWrongToken(t *testing.T) {
	err := Init()
	if nil == err {
		t.Fatal("Should return error")
	}
}

