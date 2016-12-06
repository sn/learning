package main

import "fmt"

/*
  Concurrency:
  - Do not communicate by sharing memory, instead share memory by communicating.

  Goroutines:
  - Lightweight thread
  - Managed by the Go runtime
  - Invoked by using a simple keyword `go`
  - Calling goroutines give us a challenge of figuring out when to end our main thread

  Channels:
  - Sort of like a unix pipe
  - Channels can be buffered or unbuffered
  - Its a primitive construct in the Go Language
  - Can be used in a for look with a range

*/

func myFunction1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("myFunction1() -> %d\n", i)
	}
}

func myFunction2() {
	for i := 0; i < 10; i++ {
		fmt.Printf("myFunction2() -> %d\n", i)
	}
}

func main() {

	// This is how we create a channel that deals with boolean data but you can make it work with any other types
	done := make(chan bool, 2) // We're creating a buffered channel here than can take 2 pieces of data

	// Immediately invoke the anonymous function into a goroutine

	go func() {
		myFunction1()

		done <- true // Direction of the arrow is the direction of the data (reverse assignment)
		dont <- true // We can pass this second one because we have a buffered channel
	}()

	go myFunction2() // This get's called and blocks

	// We now wait for the data to come back via the channel so the main thread can finish

	<-done
}
