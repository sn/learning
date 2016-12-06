package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name string
	age  int
}

// This function returns three values
func returnsTwo(i int) (int, int, int) {
	return i + 1, i + 2, i + 3
}

// Takes a struct
func greet(s Person) {
	fmt.Println("Hello " + s.name + " who is " + strconv.Itoa(s.age))
}

// Multiple return values
func namedReturn() (age int, name string) {
	age = 100
	name = "Jack Sparrow"
	return
}

// Takes unknown number of parameters
func variadicFunction(greeting ...string) {
	for _, v := range greeting {
		fmt.Println(v)
	}
}

func main() {
	sean := Person{name: "Sean", age: 33}

	greet(sean)

	// Named returns

	age, name := namedReturn()

	fmt.Println(age, name)

	variadicFunction("Sean", "Ninja", "Krampus", "Jack", "Sparrow")

	fmt.Println(len("this is a string"))

	x := []string{"trhis", "asdasd"}

	fmt.Println(len(x))
}
