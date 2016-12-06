package main

import "fmt"

// A map is similar to a python dict and PHP associative array

/*
  Maps:
  - Types used as keys must have the equality operator defined
  - Maps are reference types (behaves like a pointer)
  - Not thread safe
  - First-class citizens in Golang
  - Keys must be unique
*/

// We don't need to explicitly pass map by reference
func addToMap(items map[string]string, key string, val string) {
	items[key] = val
}

// Ranges work with Maps
func printMap(items map[string]string) {

	fmt.Println("---")

	for k, v := range items {
		fmt.Println(k, v)
	}
}

func main() {
	/*
	   We use make() here to allocate the memory otherwise
	   the map is essentially nil: map[key_type]value_type
	*/
	languages := make(map[string]string)

	languages["en"] = "English"
	languages["ar"] = "Arabic"
	languages["fr"] = "French"

	addToMap(languages, "af", "Afrikaans")

	/*
	   Shorthand syntax for declaring a map (no need for make)
	*/

	languages2 := map[string]string{
		"en": "ENGLISH",
		"de": "GERMAN",
		"ar": "ARABIC",
		"fr": "FRENCH", // This final comma MUST be here
	}

	printMap(languages2)

	/*
	   Map Operations:
	   - Insert
	   - Update
	   - Delete
	   - Check for existence
	*/

	// Insert

	languages["de"] = "Germ0n"

	printMap(languages)

	// Update OR Add (Same syntax)

	languages["de"] = "German"

	printMap(languages)

	// Delete (using builtin - https://golang.org/pkg/builtin/#delete)

	delete(languages, "en")

	printMap(languages)

	// Checking for key existence

	if val, exists := languages["ar"]; exists {
		fmt.Println("The key exists:", k, v)
	} else {
		fmt.Println("Does not exist:", v)
	}
}
