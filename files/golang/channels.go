package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNumber = 0
var pizzaName = ""

func makeDough(stringChan chan string) {
	pizzaNumber++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNumber)

	fmt.Println("Make dough and send for sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addSauce(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add sauce and send", pizza, "for toppings")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addToppings(stringChan chan string) {
	pizza := <-stringChan

	fmt.Println("Add toppings to", pizza, "and ship")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func main() {

	stringChan := make(chan string)

	for i := 0; i < 5; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}
}
