package main

import "fmt"

func main() {
	x := make(map[string]string)

	x["name"] = "Shahid"
	x["age"] = "12"
	x["height"] = "5.7"
	delete(x, "age")
	fmt.Println(x)
}
