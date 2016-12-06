package main

import (
	"fmt"
	"strconv"
)

func main() {
	randInt := 1
	randFloat := 20.20
	randString := "This is my string"
	randString2 := "100"
	randString3 := "200.22"

	p(float64(randInt))
	p(int(randFloat))

	newInt, _ := strconv.ParseInt(randString, 0, 64)
	p(newInt)

	newInt2, _ := strconv.ParseInt(randString2, 0, 64)
	p(newInt2)

	newFloat, _ := strconv.ParseFloat(randString3, 64)
	p(newFloat)
}

func p(v interface{}) {
	fmt.Println(v)
}
