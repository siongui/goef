package goef

import (
	"os"
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

	err = GenerateGoPackage(&pkgdata, os.Getenv("PKGDIR"))
	if err != nil {
		t.Error(err)
		return
	}
}
