# Elixir

## Learning Resources:

- Programming Elixir 1.2

## What is it?

- It's a functional programming language inspired by the Ruby syntax that runs on the Erlang VM
- It takes the actor-based approach to concurrency
- Everything is immutable state
- Elixir is about transforming data using functions and not representing objects
- It's great for writing highly scalable and fault tolerant applications

## Syntax Notes / Guidelines

- Use 2 spaces instead of tabs
- Identifiers may use lower case, uppercase, numbers and !
- Module names need to start with a capital letter and use bumpy case
- If the first letter is an underscore, Elixir wont report it when unused
- `nil` is treated as false and all other values are treated as true
- The basic unit of scoping is the function body
- Variables defined in a module are not accessible to functions in that module
- String interpolation is the same as it is in Ruby

A style guide can be found here: https://github.com/levionessa/elixir_style_guide

## Working with files

- There are two types of files: `.ex` (for compiled files) and `.exs` (for scripting files)
- File names make use of the `snake_case` structure

## IEx

IEx is the Elixir REPL that allows you to type code in and evaluate it, its the
powerful Swiss Army knife of your Elixir applications.

You can use it to recall the documentation of functions using `h()`:

    iex> h String.downcase
    iex> h String
    iex> h Enum

IEx can be used to compile files into your active session using:

    iex> c "./path/to/file.exs"

You can import an external file into your IEx session by using:

    iex> import_file "./path/to/file.exs"

There are numerous other helper functions available in IEx, check them out:

    iex> h

You can also call some OS level applications from IEx:

    iex> ls

## Pattern Matching

In conventional languages we get assignment (=), but in Elixir you get pattern
matching or more simply put, the assertion that the left side of the equal side
matches the right hand side. The equal sign in Elixir is know as a **match operator**.

Some examples that all evaluate:

    iex> a = 1
    iex> 1 = a
    iex> b = [1,2,3,4]
    iex> [c,d,e,f] = b
    iex> list = [1,2,3]
    iex> [a,2,b] = list

You could use the underscore `_` to ignore a value in the match operation:

    iex> list = [1,2,3,4]
    iex> [a,_,c,_] = list

If you want to force Elixir to use an existing value for the variable in the match
operation, you can prefix it with the `^` caret, aka **the pin operator**.

    iex> a = 1
    iex> a = 2
    iex> ^a = 1 (Will cause an error)

You can think of the equal sign in Elixir in a similar way you would use it in
algebra:

    y = 10 + c

In the above example, you are not assigning 10 + c to y. You are asserting that
the values on either side of the match operator are equal.

## Data Immutability

The word immutable means that something cannot change once created and Elixir
enforces this onto data. In Elixir, all values are immutable.

Each operation on the data creates, modifies and returns a copy of the original
value. It can be said that we `transform` the data between functions.

Elixir (Erlang) is internally optimized to ensure that handling data this way is
not a performance issue. The garbage collector runs in the smaller isolated HEAP
spaces of each application and is therefore much faster.

The important thing to remember is that any function that transforms data will
return a new copy of it. Functional languages ALWAYS return copies of the original
data.

## Native Data Types

Elixir supports the following built-in types:

**Value Types**
- Integers
- Floats
- Atoms
- Ranges
- Regular Expressions

**System Types**
- PID's
- Ports
- References

**Collection Types**
- Tuples
- Lists
- Maps
- Binaries

Functions are types too in Elixir.

* There is no internal limit on the size of integers as their internal representation grows as needed.
* Atoms names are their value, similar to Ruby Symbols.
* Atoms with the same name will always compare as equal despite being create in vastly different locations (even servers)
* Elixir has regular expression literals
* A list in Elixir is a linked data structure - it must be empty or consist of a `head` and `tail`.
* Lists can be expensive when trying to access an element at a specific position.

## Keyword Lists

Using a combination of lists of tuples, Elixir gives us a shortcut for creating lists with
simple key/value pairs.

The following are equivalent:

    [{first_name: "Sean"}, {last_name: "Nieuwoudt"}]
    [first_name: "Sean", last_name: "Nieuwoudt"]
    DB.save(record, first_name: "Sean", last_name: "Nieuwoudt")

It's important to note that a key can be repeated in a keyword list.

## Maps

A map is a collection of key => value pairs:

    iex> map = %{key => value, key => value}

In a map, keys are expected to be unique. A map can be thought of as being
similar to an associative array in other languages.

You extract values from a map using the key and bracket syntax:

    iex> map = %{key => value}
    iex> IO.puts map[key]

If the keys are `:atoms`, you can use the dot-notation to access values in the map:

    iex> map = %{:key => value}
    iex> map.key

## Anonymous Functions

Since Elixir is a functional language, a function is a basic type. Functions are
at the core of Elixir and it's data transformation power.

You can create an anonymous function using this syntax:

```
    fn
      params -> body
    end
```

As an example:

```
iex> sum = fn (a, b) -> a + b end
iex> IO.puts sum.(1, 2)
```

To call an anonymous function, you use the variable name that the function is
assigned to appended by a period "." - this is a difference between named and
anonymous functions.

Even if your function takes no parameters, you still need to use the parenthesis
when calling it.

A single function allows you to define multiple implementations depending on the type
and content of the arguments passed

```
multiplex = fn
  ({:ok,   _}) -> IO.puts "Ok!"
  ({:fail, _}) -> IO.puts "Fail!"
  ({_,     _}) -> IO.puts "Catch All"
end
```

Functions can also return other functions, but this can become cumbersome if not
clearly specified.

## Functions Returning Functions

Functions can return other functions that retain the variable bindings from the scope
of their parent function. This is possible since functions are types.

## Passing Functions as Arguments

Since functions are just types in Elixir, we can pass them to other functions as
arguments.

This is vitally important as we use this ability throughout Elixir.

```
Enum.map [1,2,3,4], fn el -> el * 2 end
```

## The &Notation in Elixir

The `&notation` is shortcut in creating functions. These two examples are the same:

```
fun1 = &(&1 + 1)
fun2 = fn a -> a + 1 end
```

The `&` operator converts the expression that follows into a function. The
tokens &1, &2, &3 etc are converted to the parameters of the function.

You can also use the & operator to create a kind of alias to other known functions:

```
len = &length/1
len.([1,2,3,4])
```

The & notation also gives you a really nice way to pass functions as parameters
to other functions.

```
Enum.map [1,2,3,4], &(&1 * 2)
```










## Subjects Needing Further Reading

1. Elixir regular expressions
2. The `with` expression
3. Parameterized functions
