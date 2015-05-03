BUILDDIR := ${CURDIR}
GO 		:= env GOPATH="${BUILDDIR}" go

all: build

build:
	${GO} build -o ./jsonf

install: build
	cp ./jsonf /usr/bin/

clean: ./jsonf
	rm -rf ./jsonf
