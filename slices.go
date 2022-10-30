package main

import "fmt"

func main() {

	//declaring  a slice in go
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("len: ", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	c := make([]string, len(s))
	copy(c, s)
	l := s[2:5]
	fmt.Println("sl1:", l)
	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3", l)
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

}
