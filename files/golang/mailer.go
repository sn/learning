package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name   string
	Email  string
	Active bool `false`
}

type Message struct {
	Subject string
	Body    string
}

func (p Person) sendEmail(m Message) {
	time.Sleep(10000 * time.Millisecond)
}

func (p *Person) UpdateStatus() {
	p.Active = true
}

func main() {
	person := Person{
		Name: "Sean Nieuwoudt", Email: "sean@wixelhq.com",
	}
	message := Message{
		Subject: "Welcome Email", Body: "Welcome to our site!",
	}

	completed := make(chan Person)

	go func() {
		person.sendEmail(message)
		person.UpdateStatus()

		completed <- person
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println(<-completed)

	fmt.Println("DONE!")

}
