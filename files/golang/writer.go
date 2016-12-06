package main

import "fmt"

/*
   Note:
   - By simply create a function that corresponds to any interface methods on any named type,
     you are automatically implementing that interface on that type.
*/

// Person is our named type struct
type Person struct {
	Name  string
	Email string
}

// Here we implement the Write method defined in the Writer interface in Golang itself
func (person *Person) Write(p []byte) (n int, err error) {
	s := string(p)
	person.Change(s)
	n = len(s)
	err = nil
	return
}

// Change allows us to change the Name property on our Person named type
func (person *Person) Change(name string) {
	person.Name = name
}

func main() {
	person := Person{
		Name: "Sean", Email: "sean@wixelhq.com", // Watch out for the comma here
	}

	person.Change("Jack")

	fmt.Println(person)

	// We can now use our Person type in the writer context
	fmt.Fprintf(&person, "The count is %d", 10)

	fmt.Println(person)
}
