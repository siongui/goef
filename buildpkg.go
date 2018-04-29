// Package goef helps you embed file/data/assets/resources/binary directly in Go
// code. There are many tools can help you embed too, but this package tries to
// be with minimal features (files are read-only) and easy to use.
package goef

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"text/template"
)

const gofile = `package {{.PkgName}}

import (
	"encoding/hex"
)

func ReadFile(filename string) ([]byte, error) {
	return hex.DecodeString(myFile)
}
`

type pkgData struct {
	PkgName string
	Files   []pkgFile
}

type pkgFile struct {
	Name          string
	Base64Content string
}

func GenerateGoPackage(pkgname, dirpath, outputpath string) (err error) {
	fo, err := os.Create(outputpath)
	if err != nil {
		return
	}
	defer fo.Close()

	pd := pkgData{PkgName: pkgname}
	err = filepath.Walk(dirpath, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.Mode().IsRegular() {
			println(path)
			pd.Files = append(pd.Files,
				pkgFile{
					Name:          path,
					Base64Content: base64.StdEncoding.EncodeToString([]byte(path)),
				})
		}

		return nil
	})
	if err != nil {
		return
	}

	tmpl, err := template.New("goembed").Parse(gofile)
	if err != nil {
		return
	}

	err = tmpl.Execute(fo, pd)
	return
}
