package main

import "log"

// We can declare function types

type Logger func(string)

func out(v string, do Logger) {
	do(v)
}

func main() {
	out("asjhdakjsdhkjasdkjhas", func(s string) {
		log.Print(s)
	})
}
