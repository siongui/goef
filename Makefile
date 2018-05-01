# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

export PKGNAME=mypkg
export PKGDIR=${GOPATH}/src/github.com/siongui/${PKGNAME}

test: fmt
	rm -rf ${PKGDIR}
	mkdir -p ${PKGDIR}
	go test -v buildpkg.go buildpkg_test.go -args -pkgdir=${PKGDIR} -pkgname=${PKGNAME}
	go test -v import_test.go

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
