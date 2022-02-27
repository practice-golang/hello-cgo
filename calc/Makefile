all: build

cgo:
	gcc -o add/add.o -c add/add.cpp
	ar rcs libadder.a add/add.o

golang:
	go build

build: cgo golang

clean:
	rm -f *.out \
		add/*.o add/*.obj \
		*.a *.dll *.lib *.def *.exp \
		calc*
