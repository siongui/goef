package goef

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/siongui/goef/mytestpkg"
)

func isInArray(array []string, item string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}

func CommonImportTest(t *testing.T) {
	a1, err := ioutil.ReadFile("testdir/hello.txt")
	if err != nil {
		t.Error(err)
		return
	}
	a2, err := ioutil.ReadFile("testdir/backtick.txt")
	if err != nil {
		t.Error(err)
		return
	}
	a3, err := ioutil.ReadFile("testdir/subdir/hello2.txt")
	if err != nil {
		t.Error(err)
		return
	}

	b, err := mytestpkg.ReadFile("hello.txt")
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(a1, b) {
		t.Error("hello.txt content not correct")
		return
	}

	b, err = mytestpkg.ReadFile("backtick.txt")
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(a2, b) {
		t.Error("backtick.txt content not correct")
		return
	}

	b, err = mytestpkg.ReadFile("subdir/hello2.txt")
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(a3, b) {
		t.Error("subdir/hello2.txt content not correct")
		return
	}

	_, err = mytestpkg.ReadFile("hello3.txt")
	if err != os.ErrNotExist {
		t.Error("hello3.txt should not exit!")
		return
	}

	filenames := mytestpkg.MapKeys()
	if !isInArray(filenames, "hello.txt") {
		t.Error("hello.txt not in MapKeys")
		return
	}

	if !isInArray(filenames, "backtick.txt") {
		t.Error("backtick.txt not in MapKeys")
		return
	}

	if !isInArray(filenames, "subdir/hello2.txt") {
		t.Error("subdir/hello2.txt not in MapKeys")
		return
	}
}
