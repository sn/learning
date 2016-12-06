package main

import "fmt"

const X = "AE!"

const (
	PI       = 3.2123
	Language = "GO"
	A        = iota
	B        = iota
	C        = iota
)

const (
	D = iota
	E
	F
	G
)

func main() {
	fmt.Println(X)
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
}
