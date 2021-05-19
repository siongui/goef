# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef TRAVIS
	export GOROOT=$(realpath ../paligo/go)
	export GOPATH=$(realpath ./gopath)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

#export GO111MODULE=off
export PKGNAME=mypkg
export PKGDIR=${GOPATH}/src/siongui/${PKGNAME}

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

test: fmt
	rm -rf ${PKGDIR}/*.go
	go test -v ${ALL_GO_SOURCES} buildpkg_test.go -args -pkgdir=${PKGDIR} -pkgname=${PKGNAME}
	go test -v import_common_test.go import_test.go
	rm -rf ${PKGDIR}/*.go
	go test -v ${ALL_GO_SOURCES} buildpkg2_test.go -args -pkgdir=${PKGDIR} -pkgname=${PKGNAME}
	go test -v import_common_test.go import2_test.go
	rm -rf ${PKGDIR}/*.go
	go test -v ${ALL_GO_SOURCES} plaintextmaxsize_test.go -args -pkgdir=${PKGDIR} -pkgname=${PKGNAME}
	go test -v import_common_test.go import2_test.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go

modinit:
	go mod init github.com/siongui/goef

modtidy:
	go list -m all
	go mod tidy
