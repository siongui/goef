package goef

import (
	"testing"
)

func TestGenerateGoPackage(t *testing.T) {
	sl, err := FileToStringLiteral("testfile/hello.txt")
	if err != nil {
		t.Error(err)
		return
	}

	pkgdata := PkgData{
		PkgName:          "myvfs",
		StringLiteralHex: sl,
	}

	err = GenerateGoPackage(&pkgdata, "src/github.com/siongui/myvfs")
	if err != nil {
		t.Error(err)
		return
	}
}
