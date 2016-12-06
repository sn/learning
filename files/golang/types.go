package main

import "fmt"

/**
 * NB: Anything that has a capital letter will allow it to be visible outside of
 * package.
 */

// Salutation is the ..
type Salutation string // This is a custom type

// Box is a struct to contain...
type Box struct {
	label  string
	width  float64
	height float64
}

func (b Box) test() {
	fmt.Println("Test!")
}

func main() {

	var message Salutation = "Hello Sir!"

	fmt.Println(message)

	// Shorthand declaration
	var box1 = Box{"My Box", 100.0, 200.00}

	// This is more explicit
	var box2 = Box{label: "My Box", width: 100.0, height: 200.00}

	// Alternative declaration

	box3 := Box{}
	box3.height = 100.00
	box3.label = "Another box!"

	fmt.Println(box1)
	fmt.Println(box2)
	fmt.Println(box3)

	box1.test()
	box2.test()
	box3.test()
}
