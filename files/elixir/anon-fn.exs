IO.puts "sum.()"

sum = fn (a, b) -> a + b end

# sum = fn a, b -> a + b end

IO.puts sum.(1,2)

IO.puts "random.()"

random = fn -> Enum.random([1, 2, 3, 4, 5]) end

IO.puts random.()

swap = fn {x, y} -> {y, x} end

# swap.({1, 2}) => {2, 1}

list_concat = fn (a, b) -> a ++ b end

# list_concat.([:a, :b], [:c, :d]) => [:a, :b, :c, :d]

pair = fn {a, b} -> [a, b] end

# pair.({1, 2}) => [1, 2]

multiplex = fn
  ({:ok,   _}) -> IO.puts "Ok!"
  ({:fail, _}) -> IO.puts "Fail!"
  ({_,     _}) -> IO.puts "Catch All"
end

multiplex.({:ok, 1})
multiplex.({:fail, 1})
multiplex.({2, 2}) #> Catch All

read_file = fn
  {:ok, file} -> "Reading file: #{IO.read(file, :line)}"
  {_, error}  -> "Error: #{:file.format_error(error)}"
end

IO.puts read_file.(File.open("/Users/sean/Github/learning/files/elixir/anon-fn.exs"))
IO.puts read_file.(File.open("invalid"))

# FizzBuzz

fizz = fn
  (0, 0, _) -> "FizzBuzz"
  (0, _, _) -> "Fizz"
  (_, 0, _) -> "Buzz"
  (_, _, c) -> c
end

IO.puts fizz.(0, 0, 1)
IO.puts fizz.(0, 1, 1)
IO.puts fizz.(1, 0, 1)
IO.puts fizz.(1, 1, 1)

fb = fn
  n -> fizz.(rem(n, 3), rem(n, 5), n)
end

results = [fb.(10), fb.(11), fb.(12), fb.(13), fb.(14), fb.(15), fb.(16)]

IO.inspect results

## Returning functions

fun1 = fn -> fn -> "Hello" end end

IO.inspect fun1.().()

fun2 = fn name -> (fn -> "Hello #{name}" end) end

IO.inspect fun2.("Sean").()

## Parameterized functions

add1 = fn n -> (fn other -> "#{n} #{other}" end) end

IO.inspect add1.("Mr").("Sean")

## Passing functions as arguments

fun3 = fn n -> n * 10 end
fun4 = fn b, fun3 -> fun3.(b) end

IO.inspect fun4.(10, fun3)

## Function Variable Pinning

defmodule Greeter do
  def for(name, greeting) do
    fn
      (^name) -> "#{greeting} #{name}"
      (_) -> "Who are you?"
    end
  end
end

greeter = Greeter.for("Sean", "Hi")

IO.inspect greeter.("Sean")
IO.inspect greeter.("Unkonwn")

## The &notation

fun5 = &(&1 + 1)
fun6 = fn a -> a + 1 end

fun7 = &length/1
IO.inspect fun7.([1,2,3,4])
