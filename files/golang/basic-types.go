package main

import "fmt"

/*
   Basic Types:
   - bool
   - string
   - int, int8, int16, int32, int64
   - uint, uint8, uint16, uint32, uint64, uintptr
   - byte (uint8)
   - rune (int32), like a char
   - float32, float64
   - complex64, complex128

   Complex Types:
   - Array
   - Slice
   - Struct
   - Pointer
   - Function
   - Interface
   - Map
   - Channel
*/

func main() {

	// Arrays:
	// - Collection that has a fixed size of a particular type
	// - Fixed size
	// - Not a Pointer type, but rather a value type
	// - No initialization needed
	// - Allocated memory is zero'd out by default

	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	for i, v := range myArray {
		fmt.Println(i, v)
	}

}
