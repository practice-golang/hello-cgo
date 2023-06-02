package main // import "structs"

// #cgo CFLAGS: -I./company
// #cgo LDFLAGS: -L. -lstructs -lstdc++
// #include "company.h"
import "C"
import (
	"unsafe"
)

type StructContainer struct {
	Company unsafe.Pointer
}

func main() {
	s := StructContainer{}
	s.Company = C.init_company()

	C.generate_employee_list(s.Company, C.int(5))
	C.print_employee_list(s.Company)

	C.free_company(s.Company)
}
