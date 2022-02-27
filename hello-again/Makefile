all: build

cgo:
	gcc -o c/say.o -c c/say.c
	ar rcs lib/libmylib.a c/say.o

golang:
	go build ./go/hello-again.go

build: cgo golang

clean:
	rm -f *.out \
		c/*.o c/*.obj \
		lib/*.a lib/*.dll lib/*.lib lib/*.def lib/*.exp \
		hello-again*
