package main

import "fmt"

import "./greeting"

func main() {

	// Modules

	greeting.DoGreeting("Sean")

	// If-statements

	isBool := false

	if prefix := getPrefix("Sean"); isBool {
		fmt.Println(prefix, "is only in this scope")
	} else {
		fmt.Println("Inside the else", prefix)
	}

	// Switch Statements

}

func getPrefix(name string) (prefix string) {
	switch name {
	case "Sean":
		prefix = "Mr "
		fallthrough
	case "Bev", "Merriman":
		prefix = "Ms "
	case "Levi":
		prefix = "Dick "
	default:
		prefix = "N/A "
	}

	return
}

func getPrefix2(name string) (prefix string) {
	switch {
	case name == "Sean":
		prefix = "Mr "
		fallthrough
	case name == "Bev", name == "Merriman", len(name) > 10:
		prefix = "Ms "
	case name == "Levi":
		prefix = "Dick "
	default:
		prefix = "N/A "
	}

	return
}
