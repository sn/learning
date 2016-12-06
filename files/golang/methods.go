package main

import "fmt"

/*
	Methods in Go:
	- It's a function that specifies what type it operates on
	- Can operate on any NAMED type
	- It means you can create a method for any built in type by creating your own type from it
	- Can operate on pointers to named types
*/

// Person is our custom type of struct
type Person struct {
	ID    int
	Name  string
	Email string
}

// People is created as a named type of a Person slice
type People []Person

// Greet is where we define a method for our named type People
func (people People) Greet(greeting string) {
	for _, v := range people {
		fmt.Printf("%s %s (%s) \n", greeting, v.Name, v.Email)
	}
}

// Rename allows us to rename a variable in our Person struct
func (person *Person) Rename(newName string) {
	person.Name = newName
}

func main() {
	people := People{
		{Name: "Sean", Email: "sean@wixelhq.com", ID: 1},
		{Name: "Sean Test", Email: "sean+test@wixelhq.com", ID: 1}, // This comma MUST be here
	}

	people.Greet("Hello!")

	person := people[0]

	person.Rename("Jack")

	fmt.Println(person)

}
