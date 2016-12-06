package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	// Basic Strings

	sampleString := "This is my string"

	fmt.Println(sampleString)
	fmt.Println(strings.Contains(sampleString, "my"))
	fmt.Println(strings.Index(sampleString, "my"))
	fmt.Println(strings.Count(sampleString, "s"))
	fmt.Println(strings.Replace(sampleString, "s", "S", len(sampleString)))

	csvString := "1,2,3,4,5,6"

	fmt.Println(strings.Split(csvString, ","))

	arr := strings.Split(csvString, ",")

	for _, v := range arr {
		fmt.Println(v)
	}

	// Arrays of strings

	stringArray := []string{"s", "e", "a", "n"}

	sort.Strings(stringArray)

	fmt.Println(stringArray)

	// File IO

	file, err := os.Create("myfile.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("this is a random file")
	file.Close()

	stream, err := ioutil.ReadFile("myfile.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	fmt.Println(readString)
}
