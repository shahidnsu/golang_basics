package main

import "fmt"

func main() {

	students := []string{"Shahid", "Mainul", "Anonnya"}

	for i, std := range students {
		fmt.Println(i, std)
	}

	for _, st := range students {
		fmt.Println(st)
	}
}
