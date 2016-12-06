package main

import (
	"fmt"
	"strconv"
)

/*
  Note:
  In this example, we're going to use a goroutine to return a Person object until the thread is closed/finished
  We'll loop through the channel results with a for look and display the person name

   Select:
   - Kind of similar to a switch statement, but for channels / communications
   - Construct of the Go Language
   -

   Rules:
   1. Executes a case that is "ready"
   2. If more than one is ready, it will execute one at random
   3. If none are ready, and no default is defined, it will block until one case is ready
*/

type Person struct {
	Name string
}

type People []Person

// getPeople simply generates a list of random Person structs and pipes it to the channel
func getPeople(c chan Person) {
	for i := 0; i < 20; i++ {
		person := Person{Name: "Sean #" + strconv.Itoa(i)}
		c <- person
	}

	close(c)
}

func main() {

	// Using Selects

	c1 := make(chan Person) // Channel 1
	c2 := make(chan Person) // Channel 2

	go getPeople(c1)
	go getPeople(c2)

	for {
		select {
		case p, ok := <-c1:
			if ok {
				fmt.Println(p.Name, ":1")
			} else {
				return
			}

		case p, ok := <-c2:
			if ok {
				fmt.Println(p.Name, ":2")
			} else {
				return
			}

		default:
			fmt.Println("Waiting...")
		}
	}

}
