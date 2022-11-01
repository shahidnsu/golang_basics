package main

import "fmt"

func main() {

	fmt.Println("Enter your age: ")
	var age int
	fmt.Scanf("%d\n", &age)
	if age <= 2 {
		fmt.Println("inflant")

	} else if age > 2 && age <= 12 {
		fmt.Println("Children")

	} else if age > 12 && age <= 40 {
		fmt.Println("Adult")

	} else {
		fmt.Println("Middle aged")

	}

}
