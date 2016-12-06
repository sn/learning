package main

import "fmt"

/*
   Interface:
   - Any type that implements the method defined in any interface, automatically implements that interface
   - There's no way in Go to require that an interface be implemented
   - The power here is the ability to create functions that take an interface as a parameter which allows you to pass in multiple types to the function
   - Every single type implements the empty interface{} so it's similar to Object or StandardObject in other languages
*/

// Renameable is an interface to be implemented by our named types
type Renameable interface {
	Rename(newName string)
}

// Person if our custom struct
type Person struct {
	Name string
}

// People is a named type for slice of People
type People []Person

// Rename is our named types' method to rename the Name property
func (p *Person) Rename(newName string) {
	p.Name = newName
}

// This method accepts any type that implements the Renameable interface
func renameToNA(p Renameable) {
	p.Rename("N/A")
}

func printPeople(p People) {

	fmt.Println("---")

	for _, v := range p {
		fmt.Println(v.Name)
	}
}

func main() {
	people := People{
		{Name: "Sean"},
		{Name: "Bev"},
		{Name: "Levi"},
	}

	printPeople(people)

	person := people[0]

	renameToNA(&person) // Note that we have a de-reference operator here

	people[0] = person

	printPeople(people)
}
