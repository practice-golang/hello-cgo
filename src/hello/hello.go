package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

// MyCprint - 이럴거면 뭐하러 go를 쓰냐 C/C++을 쓰지. 뭐 필요한 곳이 있겠지...
func MyCprint(s string) {
	cs := C.CString(s)
	C.fputs(cs, (*C.FILE)(C.stdout))
	C.free(unsafe.Pointer(cs))
}

func main() {
	MyCprint("This is Print from Cgo ^-^")
}
