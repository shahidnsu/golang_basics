package main

import "fmt"

func main() {

	type Book struct {
		title  string
		author string
		isbn   int
		pages  int
		prices int
	}
	var b1 Book
	b1.title = "intro to go"
	b1.author = "colab"
	b1.isbn = 1212112
	b1.pages = 120
	b1.prices = 10
	fmt.Println(b1)

	fmt.Println(b1.title)
	fmt.Println(b1.author)
	fmt.Println(b1.isbn)
	fmt.Println(b1.prices)
	fmt.Println(b1.pages)
	//make a struct and assgin  a values

	type person struct {
		name    string
		age     int
		address string
	}
	shahid := person{
		name:    "Bithee",
		age:     12,
		address: "khilgon",
	}
	fmt.Println(shahid)
	b2 := struct {
		name    string
		age     int
		address string
	}{
		name:    "Shahid",
		age:     12,
		address: "khilgon",
	}
	fmt.Println(b2)
}
