package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Please enter your name and age: ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Your name is %s and your age is %d\n", name, age)

}
