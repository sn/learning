package main

import (
	"fmt"
	"log"
)

func perform(name string, do func(string)) {
	do(name)
}

func main() {
	action := func(n string) {
		fmt.Println(n)
	}

	perform("Sean", action)

	perform("Sean", func(n string) {
		log.Printf(n)
	})
}
