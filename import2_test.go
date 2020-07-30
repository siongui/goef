package goef

import (
	"testing"

	"github.com/siongui/mypkg"
)

func TestImport(t *testing.T) {
	CommonImportTest(t)
	filenames := mypkg.MapKeys()
	if len(filenames) != 3 {
		t.Error("number of files not correct")
		return
	}
}
