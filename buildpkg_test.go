package goef

import (
	"flag"
	"path"
	"testing"
)

var pkgdir = flag.String("pkgdir", "", "dir of package containing embedded files")
var pkgname = flag.String("pkgname", "", "dir of package containing embedded files")

func TestGenerateGoPackage(t *testing.T) {
	err := GenerateGoPackage(*pkgname, "testdir/", path.Join(*pkgdir, "data.go"))
	if err != nil {
		t.Error(err)
		return
	}
}
