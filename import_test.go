package goef

import (
	"github.com/siongui/myvfs"
	"testing"
)

func TestImport(t *testing.T) {
	b, err := myvfs.ReadMyFile()
	if err != nil {
		t.Error(err)
		return
	}

	if string(b) != "hello world\n" {
		t.Error("wrong file content")
	}

	t.Log(string(b))
}
