package main // import "calc"

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -L. -ladder
// #include "add.h"
import "C"
import "log"

func main() {
	a := C.int(3)
	b := C.int(4)

	c := C.add(a, b)
	d := uint32(c)

	log.Printf("a + b = %v\n", d)
}
