package main

import "fmt"

// normal way to delcare functions in go
func add(x int, y int) int {
	r := x + y
	return r
}

//newer way to delcare functions in go

func sum(x, y int) (res int) {
	res = x + y
	return
}

//returning mutiples values from a function
func addSquare(a, b int) (sum int, square int) {
	sum = a + b
	square = a * b
	return
}

//referencing and defrenencing on pointer
func subtract(a *int, b *int) (res int) {
	*a = *a + 10
	*b = *b + 10
	res = *a - *b
	return res
}

func main() {

	res := add(10, 20)
	fmt.Println(res)
	val := sum(10, 50)
	fmt.Println(val)
	sum, square := addSquare(2, 2)
	fmt.Println(sum, square)
	value1 := 10
	value2 := 5
	//before function call
	fmt.Println("before function call", value1, value2)

	result := subtract(&value1, &value2)

	//after function call

	fmt.Println("After function call", value1, value2)
	fmt.Println(result)
	//anoynomous function
	a := func(x, y int) (r int) {
		r = x + y
		return
	}(10, 10)
	fmt.Println(a)

}
