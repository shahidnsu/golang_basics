package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	//geting length of an array
	fmt.Println("len:", len(a))

	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)

}
