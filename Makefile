# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef TRAVIS
	export GOROOT=$(realpath ../go)
	export GOPATH=$(realpath .)
	export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)
endif

export PKGNAME=mypkg
export PKGDIR=${GOPATH}/src/github.com/siongui/${PKGNAME}

ALL_GO_SOURCES=$(shell /bin/sh -c "find *.go | grep -v _test.go")

test: fmt
	if [ -d ${PKGDIR} ]; then rm -rf ${PKGDIR}; fi
	mkdir -p ${PKGDIR}
	go test -v ${ALL_GO_SOURCES} buildpkg_test.go -args -pkgdir=${PKGDIR} -pkgname=${PKGNAME}
	go test -v import_test.go
	rm -rf ${PKGDIR}
	mkdir -p ${PKGDIR}
	go test -v ${ALL_GO_SOURCES} buildpkg2_test.go -args -pkgdir=${PKGDIR} -pkgname=${PKGNAME}
	go test -v import_test.go
	rm -rf ${PKGDIR}
	mkdir -p ${PKGDIR}
	go test -v ${ALL_GO_SOURCES} plaintextmaxsize_test.go -args -pkgdir=${PKGDIR} -pkgname=${PKGNAME}
	go test -v import_test.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
