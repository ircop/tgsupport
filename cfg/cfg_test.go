package cfg

import "testing"

func TestConfigFileDoesntExist(t *testing.T) {
	err := NewConfig("unexisting")
	if nil == err {
		t.Fatal("Should return error")
	}
}

func TestConfigFileBroken(t *testing.T) {
	err := NewConfig("testdata/empty.toml")
	if nil == err {
		t.Fatal("Should return error")
	}
}

