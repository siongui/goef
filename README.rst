==========================
Embed Files in Go_ Package
==========================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/goef?status.png
   :target: https://godoc.org/github.com/siongui/goef

.. image:: https://api.travis-ci.org/siongui/goef.png?branch=master
   :target: https://travis-ci.org/siongui/goef

.. image:: https://goreportcard.com/badge/github.com/siongui/goef
   :target: https://goreportcard.com/report/github.com/siongui/goef

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/goef/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

.. image:: https://img.shields.io/twitter/url/https/github.com/siongui/goef.svg?style=social
   :target: https://twitter.com/intent/tweet?text=Wow:&url=%5Bobject%20Object%5D

Embed Files in Go_ Package.

.. contents:: Table of Contents


Features
++++++++

- Embedded files are read-only.
- Can be used in front-end code via GopherJS_, or local Go program.
- Can be included in your Go package, or put in a separate package.
- With limit of max single file size. (only for plain-text files)


How It Works
++++++++++++

The files to be embedded in Go code will be encoded in base64_ format. The
(name, content) of the files will be put in the (key, value) pairs of Go
built-in map_ structure, and *ReadFile* method, which has the same usage as
`ioutil.ReadFile`_, is implemented for read-only access. Because base64 encoding
is used, the size of the files will increase 33%.


Install
+++++++

.. code-block:: bash

  $ go get -u github.com/siongui/goef


Usage
+++++

Assume we have following directory structure:

.. code-block:: txt

  testdir/
  ├── hello.txt
  └── subdir/
      └── hello2.txt

We want to embed the files in *testdir/* to our code, i.e., embed *hello.txt*
and *subdir/hello2.txt* in our code, and the name of our package is *mypkg*. You
can embed the files as follows:

.. code-block:: go

  package main

  import (
  	"github.com/siongui/goef"
  )

  func main() {
  	err := goef.GenerateGoPackage("mypkg", "testdir/", "data.go")
  	if err != nil {
  		panic(err)
  	}
  }

The above code will generate *data.go* in current directory, which contains the
files directly in the code. You can read embedded files with the following
method:

.. code-block:: go

  func ReadFile(filename string) ([]byte, error)

which has the same usage as `ioutil.ReadFile`_ in Go standard library. You can
read *hello.txt* as follows:

.. code-block:: go

  b, err := ReadFile("hello.txt")
  if err != nil {
  	// handle error here
  }

And read *subdir/hello2.txt* as follows:

.. code-block:: go

  b, err := ReadFile("subdir/hello2.txt")
  if err != nil {
  	// handle error here
  }

Note that for files in sub-directory, you have also include the path of sub-dir
in the filename.

If the file does not exit, *os.ErrNotExist* error will be returned.

You can also put the generated *data.go* in a separate package, import and read
embedded files in the same way.

If your files are plain texts, you can use GenerateGoPackagePlainText_ instead
of *GenerateGoPackage*. It is the same except that the file content is stored
in plain text instead of base64 format, and the size will not increase 33%
because of base64 encoding.

GenerateGoPackagePlainTextWithMaxFileSize_ is the same as
GenerateGoPackagePlainText_ except the output file size cannot be over the given
max limit. This is useful for deploy your code on cloud services such as Google
App Engine because they usually limit the max size of a single file.

For more details, see test files `buildpkg_test.go <buildpkg_test.go>`_ and
`import_test.go <import_test.go>`_.


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] | `GitHub - UnnoTed/fileb0x: simple customizable tool to embed files in go <https://github.com/UnnoTed/fileb0x>`_
       | `GitHub - jteeuwen/go-bindata: A small utility which generates Go code from any file. Useful for embedding binary data in a Go program. <https://github.com/jteeuwen/go-bindata>`_
       | `GitHub - elazarl/go-bindata-assetfs: Serves embedded files from \`jteeuwen/go-bindata\` with \`net/http\` <https://github.com/elazarl/go-bindata-assetfs>`_
       | `GitHub - GeertJohan/go.rice: go.rice is a Go package that makes working with resources such as html,js,css,images,templates, etc very easy. <https://github.com/GeertJohan/go.rice>`_
       | `GitHub - shurcooL/vfsgen: Takes an input http.FileSystem (likely at go generate time) and generates Go code that statically implements it. <https://github.com/shurcooL/vfsgen>`_
       | `GitHub - tv42/becky: Go asset embedding for use with \`go generate\` <https://github.com/tv42/becky>`_
       | `GitHub - rakyll/statik: Embed static files into a Go executable <https://github.com/rakyll/statik>`_
       | `GitHub - mjibson/esc: A simple file embedder for Go <https://github.com/mjibson/esc>`_
       | `GitHub - bouk/staticfiles: staticfiles compiles a directory of files into an embeddable .go file <https://github.com/bouk/staticfiles>`_
       | `GitHub - flazz/togo: convert any file to Go source <https://github.com/flazz/togo>`_
       | `GitHub - inconshreveable/go-update: Build self-updating Golang programs <https://github.com/inconshreveable/go-update>`_
       | `GitHub - aprice/embed: Static content embedding for Golang <https://github.com/aprice/embed>`_
       | `GitHub - gobuffalo/packr: The simple and easy way to embed static files into Go binaries. <https://github.com/gobuffalo/packr>`_

.. [2] | `Is including assets (with a tool like go-bindata) an anti-pattern? : golang <https://www.reddit.com/r/golang/comments/60166q/is_including_assets_with_a_tool_like_gobindata_an/>`_
       | `How to build Go plugin with data inside : golang <https://www.reddit.com/r/golang/comments/63f3ag/how_to_build_go_plugin_with_data_inside/>`_
       | `golang - compile static files in app? : golang <https://www.reddit.com/r/golang/comments/66uewv/golang_compile_static_files_in_app/>`_
       | `embed: Yet Another Static Content Embedder for Go : golang <https://www.reddit.com/r/golang/comments/6fh80b/embed_yet_another_static_content_embedder_for_go/>`_
       | `Embed libraries into exe : golang <https://www.reddit.com/r/golang/comments/7h9kcx/embed_libraries_into_exe/>`_

.. [3] `Embed Data in Front-end Go Code <https://siongui.github.io/2017/04/08/go-embed-data-in-frontend-code/>`_

.. [4] | `[Q&A] //go:embed draft design : golang <https://old.reddit.com/r/golang/comments/hv96ny/qa_goembed_draft_design/>`_
       | `[Q&A] io/fs draft design : golang <https://old.reddit.com/r/golang/comments/hv976o/qa_iofs_draft_design/>`_


.. _Go: https://golang.org/
.. _GopherJS: https://github.com/gopherjs/gopherjs
.. _base64: https://en.wikipedia.org/wiki/Base64
.. _map: https://blog.golang.org/go-maps-in-action
.. _ioutil.ReadFile: https://golang.org/pkg/io/ioutil/#ReadFile
.. _UNLICENSE: http://unlicense.org/
.. _GenerateGoPackagePlainText: https://godoc.org/github.com/siongui/goef#GenerateGoPackagePlainText
.. _GenerateGoPackagePlainTextWithMaxFileSize: https://godoc.org/github.com/siongui/goef#GenerateGoPackagePlainTextWithMaxFileSize
