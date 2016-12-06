package main

import "fmt"

// type User struct {
// 	Id      string
// 	Balance uint64

// }

func addThem(numbers []float64) float64 {
	sum := 0.0

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func nextTwoVals(number int) (int, int) {
	return number + 1, number + 2
}

func subtractThem(args ...int) int {

	fmt.Println("In Subtract Them")

	final := 0

	for _, v := range args {
		final -= v
	}

	return final
}

func factorial(num int) int {

	fmt.Println("In Factorial")

	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {
	// fmt.Print("Learning Go!")

	defer fmt.Printf("%d\n", factorial(3))

	subtractThem(1, 2, 3, 4, 5)
	// num3 := 3
	//
	// dbl := func() int {
	// 	num3 *= 2
	//
	// 	return num3
	// }
	//
	// fmt.Println(dbl())

	// fmt.Println(subtractThem(1, 2, 3, 4, 5, 6))

	// num1, num2 := nextTwoVals(5)
	//
	// fmt.Print(num1, num2)

	//listNums := []float64{1, 2, 3, 4, 5}

	//fmt.Printf("%f", addThem(listNums))

	//fmt.Println("AddThem: " + )

	// Maps

	// bikes := make(map[string]int)
	//
	// bikes["merida"] = 100
	// bikes["rocky mountain"] = 101
	//
	// delete(bikes, "merida")

	// mySlices := make([]int, 5, 10)
	//
	// for _, v := range mySlices {
	// 	fmt.Println(v)
	// }

	// copy()

	// var age int = 23
	// var num float64 = 1.3434

	// randNum := 1
	//
	// fmt.Println(randNum)
	//
	// randNum = 21313
	//
	// fmt.Println(randNum)

	// const pi float64 = 3.2323232
	//
	// var (
	// 	varA = 2
	// 	varB = 3
	// )
	//
	// fmt.Println(varA)
	// fmt.Println(varB)
	//
	// fmt.Println(len("hello" + "\t" + "world"))
	//
	// isAllowed := false
	//
	// fmt.Printf("%t \n", isAllowed)

	// i := 1
	//
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// u := User{Id: "US123", Balance: 8}
	// b := new(bytes.Buffer)
	// json.NewEncoder(b).Encode(u)
	// res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", b)
	// io.Copy(os.Stdout, res.Body)

	// Arrays

	// myAr[5] float64
	//
	// myAr[0] = 123.123
	// myAr[1] = 123.123
	// myAr[2] = 123.123
	// myAr[3] = 123.123
	// myAr[4] = 123.123
	//
	// favNums := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//
	// for i, val := range favNums {
	// 	fmt.Printf("%d: %f \n", i, val)
	// }

	// myNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//
	// for _, v := range myNums {
	// 	fmt.Printf("%d\n", v)
	// }
	//
	// myNums2 := myNums[3:5]
	//
	// fmt.Println("-")
	// fmt.Println(myNums[5])
	// fmt.Println("-")
	//
	// for _, v := range myNums2 {
	// 	fmt.Printf("%d\n", v)
	// }
}
