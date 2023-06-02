all: build

cpp_standalone:
	g++ -O3 -o company/company.o -c company/company.cpp

cgo:
	g++ -O3 -DNOT_LIBRARY company/company.cpp -o company.exe
	ar rcs libstructs.a company/company.o

golang:
	go build

build: cpp_standalone cgo golang

clean:
	rm -f *.out \
		company/*.o company/*.obj \
		*.a *.dll *.lib *.def *.exp \
		structs.exe \
		company.exe
