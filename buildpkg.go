package goef

import (
	"os"
	"path"
	"text/template"
)

const gofile = `package {{.PkgName}}

import (
	"encoding/hex"
)

const myFile = "{{.StringLiteralHex}}"

func ReadMyFile() ([]byte, error) {
	return hex.DecodeString(myFile)
}
`

type PkgData struct {
	PkgName          string
	StringLiteralHex string
}

func GenerateGoPackage(pkgdata *PkgData, pkgdir string) (err error) {
	fopath := path.Join(pkgdir, "data.go")
	fo, err := os.Create(fopath)
	if err != nil {
		return
	}
	defer fo.Close()

	tmpl, err := template.New("goembed").Parse(gofile)
	if err != nil {
		return
	}

	return tmpl.Execute(fo, pkgdata)
}
