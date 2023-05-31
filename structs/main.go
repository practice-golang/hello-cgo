package main // import "structs"

// #cgo CFLAGS: -I./employee
// #cgo LDFLAGS: -L. -lstructs
// #include "employee.h"
import "C"
import (
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Salary float32
}

func main() {
	cEmp := C.Employee{}

	C.get_employee(&cEmp)

	goEmp := Employee{
		ID:     int(cEmp.id),
		Name:   C.GoString(cEmp.name),
		Salary: float32(cEmp.salary),
	}

	fmt.Println(goEmp)

	C.set_employee(&cEmp, 2, C.CString("Jane Smith"), 6000.0)

	goEmp.ID = int(cEmp.id)
	goEmp.Name = C.GoString(cEmp.name)
	goEmp.Salary = float32(cEmp.salary)

	fmt.Println(goEmp)

	C.free_employee(&cEmp)
}
