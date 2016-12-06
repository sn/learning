package main

import "fmt"

func safeDivide(num1 int, num2 int) int {
	defer func() {
		recover()
	}()

	num := num1 / num2

	return num
}

func showPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("OUCH!")
}

func main() {
	// fmt.Println(saveDivide(1, 2))
	// fmt.Println(safeDivide(1, 0))

	showPanic()
}
