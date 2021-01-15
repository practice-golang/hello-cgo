package main // import "hello-again"

// #cgo CFLAGS: -I../include
// #cgo LDFLAGS: -L../lib -lmylib
// #include <hello.h>
// #include <sum.h>
import "C"
import (
	"errors"
	"fmt"
	"log"
)

func sayHello() {
	_, err := C.SayHello()
	if err != nil {
		log.Fatal(err)
	}
}

func doSum(a, b int) (int, error) {
	//Convert Go ints to C ints
	ac := C.int(a)
	bc := C.int(b)

	cc, err := C.add(ac, bc)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	res := int(cc)

	return res, nil
}

func main() {
	sayHello()

	res, err := doSum(1, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n1 + 2 =", res)
}
