package main

import (
	"fmt"
	"time"
)

func count(id int) {
	for i := 0; i < 10; i++ {

		fmt.Printf("%d : %d\n", id, i)

		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {

	for x := 0; x < 10; x++ {
		go count(x)
	}

	time.Sleep(time.Millisecond * 11000)
}
