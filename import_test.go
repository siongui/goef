package goef

import (
	"os"
	"testing"

	"github.com/siongui/mypkg"
)

func TestImport(t *testing.T) {
	b, err := mypkg.ReadFile("hello.txt")
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != "hello world\n" {
		t.Error("hello.txt content not correct")
		return
	}

	b, err = mypkg.ReadFile("subdir/hello2.txt")
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != "hello world2\n" {
		t.Error("subdir/hello2.txt content not correct")
		return
	}

	b, err = mypkg.ReadFile("hello3.txt")
	if err != os.ErrNotExist {
		t.Error("hello3.txt should not exit!")
		return
	}
}
