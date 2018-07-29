package logger

import "testing"

func TestEmptyLogfile(t *testing.T) {
	err := SetPath("")
	if nil != err {
		t.Fatal("Should not return error")
	}
}

func TestWrongPath(t *testing.T) {
	err := SetPath("/the/wrong/path")
	if nil == err {
		t.Fatal("Should return error")
	}
}

func TestCorrectPath(t *testing.T) {
	err := SetPath("testdata/test.log")
	if nil != err {
		t.Fatal("Should not return error")
	}
}