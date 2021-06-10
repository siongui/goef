package goef

import (
	"testing"

	"github.com/siongui/goef/mytestpkg"
)

func TestImport(t *testing.T) {
	CommonImportTest(t)
	filenames := mytestpkg.MapKeys()
	if len(filenames) != 3 {
		t.Error("number of files not correct")
		return
	}
}
