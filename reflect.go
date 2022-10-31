package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := make([]string, 0)
	fmt.Println(x)
	b := reflect.TypeOf(x).Kind().String()
	fmt.Print(b, "\n")

}
