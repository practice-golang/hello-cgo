package main // import "calc"

// #cgo CFLAGS: -I./add -I./random_float -I./random_int
// #cgo LDFLAGS: -L. -lcalc
// #include <stdlib.h>
// #include "add.h"
// #include "random_float.h"
// #include "random_int.h"
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

func add() {
	a := C.int(3)
	b := C.int(4)

	c := C.add(a, b)
	d := uint32(c)

	log.Printf("a + b = %v\n", d)
}

func randomeFloat1() {
	floatArrayLength := 5

	cFloatArrayPtr := C.random_float_pointer_array(C.int(floatArrayLength))
	defer C.free(unsafe.Pointer(cFloatArrayPtr))

	// /* float32 */
	// floatArray := make([]float32, floatArrayLength)
	// for i := 0; i < floatArrayLength; i++ {
	// 	floatArray[i] = float32(*(*C.float)(unsafe.Pointer(uintptr(unsafe.Pointer(cFloatArrayPtr)) + uintptr(i)*unsafe.Sizeof(C.float(0)))))
	// }

	/* float64 */
	floatArray := make([]float64, floatArrayLength)
	for i := 0; i < floatArrayLength; i++ {
		floatArray[i] = float64(*(*C.float)(unsafe.Pointer(uintptr(unsafe.Pointer(cFloatArrayPtr)) + uintptr(i)*unsafe.Sizeof(C.float(0)))))
	}

	for _, val := range floatArray {
		fmt.Println(val)
	}
}

func randomeFloat2() {
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

func randomeInt1() {
	intArrayLength := 5

	cIntArrayPtr := C.random_int_pointer_array(C.int(intArrayLength))
	defer C.free(unsafe.Pointer(cIntArrayPtr))

	/* int */
	intArray := make([]int, intArrayLength)
	for i := 0; i < intArrayLength; i++ {
		intArray[i] = int(*(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(cIntArrayPtr)) + uintptr(i)*unsafe.Sizeof(C.int(0)))))
	}

	// /* int32 */
	// intArray := make([]int32, intArrayLength)
	// for i := 0; i < intArrayLength; i++ {
	// 	intArray[i] = int32(*(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(cIntArrayPtr)) + uintptr(i)*unsafe.Sizeof(C.int(0)))))
	// }

	// /* float64 */
	// floatArray := make([]float64, floatArrayLength)
	// for i := 0; i < floatArrayLength; i++ {
	// 	floatArray[i] = float64(*(*C.float)(unsafe.Pointer(uintptr(unsafe.Pointer(cFloatArrayPtr)) + uintptr(i)*unsafe.Sizeof(C.float(0)))))
	// }

	for _, val := range intArray {
		fmt.Println(val)
	}
}

func randomeInt2() {
	intArrayLength2 := C.int(10)
	intArray2 := make([]int32, intArrayLength2)

	fmt.Println("intArrayLength2 #1:", intArrayLength2)

	C.random_int_by_pointer_set((*C.int)(unsafe.Pointer(&intArray2[0])), &intArrayLength2)

	fmt.Println("floatArrayLength2 #2:", intArrayLength2)

	resizeLength := int(intArrayLength2)
	intArray2 = intArray2[:resizeLength:resizeLength]

	fmt.Println("intArrayLength2:", intArrayLength2)
	fmt.Println("intArray2:", intArray2)

	for _, val := range intArray2 {
		fmt.Println(val)
	}
}

func main() {
	// add() // Add

	// randomeFloat1() // Random float #1
	// randomeFloat2() // Random float #2

	randomeInt1() // Random int #1
	// randomeInt2() // Random int #2
}
