package main

import "fmt"

/*
  Slices:
  - Maps are reference types (behaves like a pointer)
  - Very similar to vectors or arrays in other languages
  - Have flexibility of the more advanced data types
  - Memory and performance characteristics of simpler data types
  - Can be thought of as an abstraction over an array
  - The size of a slice does not effect its type
  - Fixed size at first but can be reallocted with `append()`
  - Use `make()` to initialize otherwise it will be nil
  - Points to an array underneath
*/

func printSlice(slice []int) {

	fmt.Println("---")

	for k, v := range slice {
		fmt.Println(k, v)
	}
}

func main() {

	// Initial length and capacity of 3
	a := make([]int, 3)

	a[0] = 1
	a[1] = 2
	a[2] = 3

	// Assigning past the intial capacity cause an error
	// s[3] = 4

	printSlice(a)

	// We could also directly initialize and assign values

	b := []int{1, 2, 3, 4}

	printSlice(b)

	// Extracting a slice of a slice

	printSlice(b[1:2])
	printSlice(b[2:])
	printSlice(b[:3])
	printSlice(b[0:3])
	printSlice(b[:])

	// Appending to a slice

	b = append(b, 5)

	// Here we are using `...` to unpack the slice a
	// and add each element to the b slice

	b = append(b, a...)

	printSlice(b)

	// Deleting from a slice
	// This is non-intuitive as you need to create a new slice from appending to other slices together
	// skipping over the elements from the original slice

	newSlice := append(a[:1], b[:1]...)

	printSlice(newSlice)
}
