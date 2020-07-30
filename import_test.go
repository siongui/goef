package goef

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/siongui/mypkg"
)

func TestImport(t *testing.T) {
	CommonImportTest(t)

	sl, err := ioutil.ReadFile("testdir/subdir/testlink")
	if err != nil {
		t.Error(err)
		return
	}

	sl2, err := ioutil.ReadFile("testdir/testlink2")
	if err != nil {
		t.Error(err)
		return
	}

	b, err := mypkg.ReadFile("subdir/testlink")
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(sl, b) {
		t.Error("subdir/testlink content not correct")
		return
	}

	b, err = mypkg.ReadFile("testlink2")
	if err != nil {
		t.Error(err)
		return
	}
	if !bytes.Equal(sl2, b) {
		t.Error("testlink2 content not correct")
		return
	}

	filenames := mypkg.MapKeys()
	if len(filenames) != 5 {
		t.Error("number of files not correct")
		return
	}

	if !isInArray(filenames, "subdir/testlink") {
		t.Error("subdir/testlink not in MapKeys")
		return
	}

	if !isInArray(filenames, "testlink2") {
		t.Error("testlink2 not in MapKeys")
		return
	}
}
