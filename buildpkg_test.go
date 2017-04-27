package goef

import (
	"flag"
	"testing"
)

var pkgdir = flag.String("pkgdir", "", "dir of package containing embedded files")

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

	err = GenerateGoPackage(&pkgdata, *pkgdir)
	if err != nil {
		t.Error(err)
		return
	}
}
