package main // import "hello"

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// MyCprint - Call C function
func MyCprint(s string) {
	cs := C.CString(s)
	C.fputs(cs, (*C.FILE)(C.stdout))
	C.free(unsafe.Pointer(cs))
}

func main() {
	MyCprint("This is Print from Cgo ^-^")
}
