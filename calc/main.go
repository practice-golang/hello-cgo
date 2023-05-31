package main // import "calc"

// #cgo CFLAGS: -I./add -I./random_float
// #cgo LDFLAGS: -L. -lcalc
// #include <stdlib.h>
// #include "add.h"
// #include "random_float.h"
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

func main() {
	a := C.int(3)
	b := C.int(4)

	c := C.add(a, b)
	d := uint32(c)

	log.Printf("a + b = %v\n", d)

	floatArrayLength := 10

	cFloatArrayPtr := C.random_float_pointer_array(C.int(floatArrayLength))
	defer C.free(unsafe.Pointer(cFloatArrayPtr))

	/* float32 */
	// floatArray := (*[1 << 30]float32)(unsafe.Pointer(cFloatArrayPtr))[:floatArrayLength:floatArrayLength]
	floatArray := make([]float32, floatArrayLength)
	for i := 0; i < floatArrayLength; i++ {
		floatArray[i] = float32(*(*C.float)(unsafe.Pointer(uintptr(unsafe.Pointer(cFloatArrayPtr)) + uintptr(i)*unsafe.Sizeof(C.float(0)))))
	}

	/* float64 */
	// floatArray := make([]float64, floatArrayLength)
	// for i := 0; i < floatArrayLength; i++ {
	// 	floatArray[i] = float64(*(*C.float)(unsafe.Pointer(uintptr(unsafe.Pointer(cFloatArrayPtr)) + uintptr(i)*unsafe.Sizeof(C.float(0)))))
	// }

	for _, val := range floatArray {
		fmt.Println(val)
	}

	floatArrayLength2 := C.int(10)
	floatArray2 := make([]float32, floatArrayLength2)

	fmt.Println("floatArrayLength2 #1:", floatArrayLength2)

	C.random_float_by_pointer_set((*C.float)(unsafe.Pointer(&floatArray2[0])), &floatArrayLength2)

	fmt.Println("floatArrayLength2 #2:", floatArrayLength2)

	resizeLength := int(floatArrayLength2)
	floatArray2 = floatArray2[:resizeLength:resizeLength]

	fmt.Println("floatArrayLength2:", floatArrayLength2)
	fmt.Println("floatArray2:", floatArray2)

	for _, val := range floatArray2 {
		fmt.Println(val)
	}
}
