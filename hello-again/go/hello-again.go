package main // import "hello-again"

// #include <say.h>
// #cgo CFLAGS: -I../include
// #cgo LDFLAGS: -L../lib -lmylib
import "C"
import (
	"log"
)

func sayHello() {
	_, err := C.SayHello()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	sayHello()
}
