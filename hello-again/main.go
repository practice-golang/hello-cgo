package main // import "hello-again"

// #cgo CFLAGS: -I./include
// #cgo LDFLAGS: -L./lib -lmylib
// #include <say.h>
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
