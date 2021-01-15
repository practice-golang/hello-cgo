all: build

cgo:
	gcc -o c/say.o -c c/say.c
	gcc -o c/sum.o -c c/sum.c
	ar rcs lib/libmylib.a c/say.o c/sum.o

golang:
	go build ./go/hello-again.go

build: cgo golang

clean:
	rm -f c/*.o lib/*.a hello-again*
