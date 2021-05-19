package goef

import (
	"testing"

	"siongui/mypkg"
)

func TestImport(t *testing.T) {
	CommonImportTest(t)
	filenames := mypkg.MapKeys()
	if len(filenames) != 3 {
		t.Error("number of files not correct")
		return
	}
}
