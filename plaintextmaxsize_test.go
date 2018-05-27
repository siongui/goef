package goef

import (
	"flag"
	"testing"
)

var pkgdir = flag.String("pkgdir", "", "dir of package containing embedded files")
var pkgname = flag.String("pkgname", "", "dir of package containing embedded files")

func TestGenerateGoPackagePlainTextWithMaxFileSize(t *testing.T) {
	err := GenerateGoPackagePlainTextWithMaxFileSize(*pkgname, "testdir/", *pkgdir, 200)
	if err != nil {
		t.Error(err)
		return
	}
}
