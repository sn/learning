package main

import "fmt"

// Rectangle is the representaion of an
type Rectangle struct {
	left   float64
	right  float64
	top    float64
	bottom float64
}

// Add a method to a struct
func (rect *Rectangle) area() float64 {
	return rect.left * rect.right
}

func main() {

	rect := Rectangle{left: 100.00, right: 20.00, top: 23.11, bottom: 30.00}

	fmt.Printf("The rectangle is %f wide \n", rect.bottom)

	fmt.Println(rect.area())

}
