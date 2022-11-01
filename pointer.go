package main

import "fmt"

func main() {

	var a int
	a = 10
	var b *int
	b = &a
	fmt.Println("a value is ", a)
	fmt.Println("b value is ", &b)
	//deferencing to get the value content by b
	fmt.Println("b contents the value", *b)
}
