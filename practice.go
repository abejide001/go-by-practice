package main

import (
	"fmt"
)

func main() {
	slice()
	mapping()
	obj()
}

// using slice
func slice() {
	y := []string{"money", "cash", "adsd", "dsfsf", "fdgggg"}
	x := []string{"money", "cash", "adsd", "dsfsf", "fdgggg", "ghost"}
	rem := [][]string{x, y}
	fmt.Println(rem)
}

// mapping
func mapping() {
	x := map[string]int{
		"John": 90,
	}
	x["Mumum"] = 20
	delete(x, "Mumum")
	fmt.Println(x)
}

// Person struct
type Person struct {
	firstName string
	lastName  string
}

func obj() {
	fullName := Person{
		firstName: "john",
		lastName:  "aliu",
	}
	fmt.Println("full name is", fullName)
}
