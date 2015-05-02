BUILDDIR := ${CURDIR}
GO 		:= env GOPATH="${BUILDDIR}" go

all: build

build:
	${GO} build -o ./jsonf

clean: ./jsonf
	rm -rf jsonf
