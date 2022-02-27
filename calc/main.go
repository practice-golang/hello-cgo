package main // import "calc"

// __declspec(dllimport) int add(int,int);
// #cgo LDFLAGS: -L. -lmylib
import "C"

import "fmt"

func main() {
	a := C.int(3)
	b := C.int(4)
	c := C.add(a, b)
	d := uint32(c)
	fmt.Printf("a*b = %v\n", d)
}
