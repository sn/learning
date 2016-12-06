package main

import "fmt"

func changesACopy(val int) int {
	val++

	return val
}

func changesAPointer(val *int) {
	*val++
}
func main() {

	originalVal := 100

	// originalValPtr := &originalVal

	fmt.Println(originalVal)

	changesACopy(originalVal)

	fmt.Println(originalVal)

	changesAPointer(&originalVal)

	fmt.Println(originalVal)
}
