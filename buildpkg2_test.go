package goef

import (
	"flag"
	"path"
	"testing"
)

var pkgdir = flag.String("pkgdir", "", "dir of package containing embedded files")
var pkgname = flag.String("pkgname", "", "dir of package containing embedded files")

func TestGenerateGoPackagePlainText(t *testing.T) {
	err := GenerateGoPackagePlainText(*pkgname, "testdir/", path.Join(*pkgdir, "data.go"))
	if err != nil {
		t.Error(err)
		return
	}
}
