package main

import "fmt"

func main() {

	// Basic for loop

	counter := 10

	for i := 0; i < counter; i++ {
		fmt.Println(i)
	}

	// While loop type usage

	i := 0

	for i < counter {
		fmt.Println(i)
		i++
	}

	// Infinite Loop

	for {
		// fmt.Println("Inside the loop")
	}
}
