package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please enter your name.")
	var name string
	//	The var name string line will create a new variable using the var keyword. You name the variable name, and it will be of type string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("Hi, %s! I'm Go!", name)

}
