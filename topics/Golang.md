# Golang

## Installation

@TODO ENV vars
@TPDP Gopaths

## General Notes

Go is a compiled language that compiles to machine code and needs no further dependencies to run. It's syntax uses unicode.

Go code is organized into packages that tend to be similar to modules in other languages.
The `package main` directive is special and indicates a standalone application, not a library.  The function `main()` is also special in `package main` since this is where processing will begin.

When importing packages, you must only import the ones you need otherwise your application will not compile. This is part of the golang spec and ensures that you only include what you need.

Go does not require semi-colons after statements or declarations unless you use two on the same line straight after each other. Go takes a very strong stance on syntax formatting and you must adhere to these rules for consistency.

The `for loop` is the only looping statement in go and it has a few different constructs.

The `blank identifier`, whose name is `_` is used when you need to assign a variable value and not use it. This will avoid compilation errors thrown by Go due to unused variables.

## Command Line

go build main.go (creates the binary)
go run main.go
gofmt main.go
go get <package>
go fmt
go <shows all commands>
go install
go clean <removes binary>
go get -u <package> (updates package)

## Data Types

There are numerous ways to declare variables in Go:

	s := "" // Not used for package level variables
    var s string
    var s = ""
    var s string = ""

## Packages

When a function name does not start with a capital letter, it is not exported outside of the package itself. It will only be availble to call from other functions inside hte package. Generally though, you don't mention public / private visibility in golang.
fmt
os

# Learning Resources


# Functions

Firstly, you can only have one `func main()`. Only functions that start with a capital letter will be exported if its part of a module.

func identifier(name <type> ...) <return_type> {
	return <return_type>
}






# Go programming

Video Link: https://www.youtube.com/watch?v=CF9S4QZuV30

* Every go program starts with a package
* Variables in Go are statically typed (they can't change once defined)
* Variables can contain letters, numbers and underscores
* There are 3 logical operators in Go ( &&, ||, !&&)
* A slice is like an array, but when defined, you leave out the size

You can view information about any functions and operators in Golang by using:

> godoc fmt Println



##### 3 Tenants of Golang:

- Compiled
- Garbage Collected
- Concurrent

The reason we use the structure of the workspace like we do is because it can do
away with build files entirely so you don't need a build management system.

Types in Golang, even from the same grouping, are not compatible in Golang. You
will need to cast them between types if you want to use them that way.

Types:

- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64, uintptr
- byte (uint8)
- rune (int32), like a char
- float32, float64
- complex64, complex128
- Array
- Slice
- Struct
- Pointer
- Function
- Interface
- Map
- Channel

* If you don't specify the types, Go will infer the type from the contents.

#### Switch Statements

- Switch statements in go have no default fall through
- You can also switch on the type of a variable, not just the value

#### Ranges

Ranges can work with the following types:

- Array or Slice
- Strings
- Maps
- Channels

#### Arrays or Slices
