package goef

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const gofile2 = `package {{.PkgName}}

import (
	"os"
)

var virtualFilesystem = map[string]string{
{{ range .Files }}"{{ .Name }}": ` + "`" + `{{ .PlainTextContent }}` + "`" + `,
{{ end }}}

func ReadFile(filename string) ([]byte, error) {
	content, ok := virtualFilesystem[filename]
	if ok {
		return []byte(content), nil
	}
	return nil, os.ErrNotExist
}
`

type pkgData2 struct {
	PkgName string
	Files   []pkgFile2
}

type pkgFile2 struct {
	Name             string
	PlainTextContent string
}

func getPlainTextContent(dirpath, path string, info os.FileInfo) (name, content string, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	content = string(b)
	// escape backtick `
	content = strings.Replace(content, "`", "`"+`+"`+"`"+`"+`+"`", -1)
	name, err = filepath.Rel(dirpath, path)
	return
}

// This method is the same as *GenerateGoPackage* method, except that file
// content are embedded directly in the code without encoding the content in
// base64 format. If your file content are plain texts, use this method instead
// of *GenerateGoPackage* method.
func GenerateGoPackagePlainText(pkgname, dirpath, outputpath string) (err error) {
	fo, err := os.Create(outputpath)
	if err != nil {
		return
	}
	defer fo.Close()

	pd := pkgData2{PkgName: pkgname}
	err = filepath.Walk(dirpath, func(filepath string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.Mode().IsRegular() {
			name, content, errf := getPlainTextContent(dirpath, filepath, info)
			if errf != nil {
				return errf
			}

			pd.Files = append(pd.Files,
				pkgFile2{
					Name:             name,
					PlainTextContent: content,
				})
		}

		return nil
	})
	if err != nil {
		return
	}

	tmpl, err := template.New("goembed").Parse(gofile2)
	if err != nil {
		return
	}

	err = tmpl.Execute(fo, pd)
	return
}
