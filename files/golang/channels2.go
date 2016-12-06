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

// Generate uses the name type value to create the range of Person structs and return it to the channel
func (people People) Generate(c chan Person) {
	for _, v := range people {
		c <- v
	}

	close(c)
}

// getPeople simply generates a list of random Person structs and pipes it to the channel
func getPeople(c chan Person) {
	for i := 0; i < 20; i++ {
		person := Person{Name: "Sean #" + strconv.Itoa(i)}
		c <- person
	}

	close(c)
}

func main() {
	c := make(chan Person) // Person is the data type that can be passed to and from this channel

	people := People{
		{Name: "Sean"},
		{Name: "Bev"},
	}

	go people.Generate(c)

	for s := range c {
		fmt.Println(s.Name)
	}

	// Alternative Example

	p := make(chan Person) // Create another channel for the alternative example

	go getPeople(p)

	for s := range p {
		fmt.Println(s.Name)
	}
}
