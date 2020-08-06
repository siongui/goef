package goef

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

const datafile = `package {{.PkgName}}

// DO NOT EDIT. This file is auto-created by github.com/siongui/goef

var vfs{{.N}} = map[string]string{
{{ range .Files }}"{{ .Name }}": ` + "`" + `{{ .PlainTextContent }}` + "`" + `,
{{ end }}}
`

const readfile = `package %s

// DO NOT EDIT. This file is auto-created by github.com/siongui/goef

import (
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	%s
	return nil, os.ErrNotExist
}

func MapKeys() []string {
	length := 0
	%s
	keys := make([]string, length)
	i := 0
	%s
	return keys
}
`

const readcode = `
	content%d, ok%d := vfs%d[filename]
	if ok%d {
		return []byte(content%d), nil
	}
`

const lengthcode = `
	length += len(vfs%d)
`

const appendcode = `
	for k := range vfs%d {
		keys[i] = k
		i++
	}
`

var minSize = int64(len([]byte(datafile)))

type pkgData3 struct {
	PkgName string
	N       int64
	Files   []pkgFile2
}

func createReadFile(pkgname, outputdir string, totalNumberOfDataFiles int64) (err error) {
	rd := ""
	ln := ""
	ap := ""
	for i := int64(0); i < totalNumberOfDataFiles; i++ {
		rd += fmt.Sprintf(readcode, i, i, i, i, i)
		ln += fmt.Sprintf(lengthcode, i)
		ap += fmt.Sprintf(appendcode, i)
	}
	fc := fmt.Sprintf(readfile, pkgname, rd, ln, ap)

	outputpath := path.Join(outputdir, "read.go")
	err = ioutil.WriteFile(outputpath, []byte(fc), 0644)
	return
}

func createDataFile(pd pkgData3, outputdir string) (err error) {
	filename := fmt.Sprintf("data%d.go", pd.N)
	outputpath := path.Join(outputdir, filename)
	fo, err := os.Create(outputpath)
	if err != nil {
		return
	}
	defer fo.Close()

	tmpl, err := template.New("goembed").Parse(datafile)
	if err != nil {
		return
	}

	err = tmpl.Execute(fo, pd)
	return
}

// GenerateGoPackagePlainTextWithMaxFileSize is the same as
// GenerateGoPackagePlainText method, except the output file size cannot be over
// the given max limit. This is useful for deploy your code on cloud services
// such as Google App Engine because they usually limit the max size of a single
// file.
func GenerateGoPackagePlainTextWithMaxFileSize(pkgname, dirpath, outputdir string, maxSize int64) (err error) {
	if maxSize < minSize {
		return fmt.Errorf("maxSize cannot be less than %d", minSize)
	}

	// check if outputdir exists
	info, err := os.Stat(outputdir)
	if err != nil {
		if os.IsNotExist(err) {
			// outputdir does not exist, create it.
			err = os.MkdirAll(outputdir, 0755)
			if err != nil {
				return
			}
		} else {
			return
		}
	} else {
		// check if outputdir is really a dir
		if !info.IsDir() {
			return errors.New("argument *outputdir* is not a directory!")
		}
	}

	pd := pkgData3{PkgName: pkgname, N: 0}
	currentDataFileSize := minSize
	err = filepath.Walk(dirpath, func(filepath string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		if info.Mode().IsRegular() {
			name, content, errf := getFilenameAndPlainTextContent(dirpath, filepath, info)
			if errf != nil {
				return errf
			}

			// 10 includes rough estimate for "`... in template rendering
			increasedSize := int64(len([]byte(name+content))) + 10

			if increasedSize > (maxSize - minSize) {
				return fmt.Errorf("file %s is too big (Please increase maxSize %d)", name, maxSize)
			}

			currentDataFileSize += increasedSize
			if currentDataFileSize > maxSize {
				if len(pd.Files) == 0 {
					return errors.New("Please increase maxSize value")
				}

				if err3 := createDataFile(pd, outputdir); err3 != nil {
					return err3
				}
				pd = pkgData3{PkgName: pkgname, N: pd.N + 1}
				currentDataFileSize = minSize + increasedSize
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

	totalNumberOfDataFiles := pd.N
	if len(pd.Files) > 0 {
		err = createDataFile(pd, outputdir)
		if err != nil {
			return err
		}
		totalNumberOfDataFiles += 1
	}
	err = createReadFile(pkgname, outputdir, totalNumberOfDataFiles)
	return
}
