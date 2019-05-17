package main

import (
	"fmt"
)

func main() {
	slice()
	mapping()
	obj()
   // anonymous struct
	anon := struct{
		color string
	} {
		color: "black",
	}
	fmt.Println(anon.color)
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
	age       float64
}

// Agent struct
type Agent struct {
	Person
	level string
}
func obj() {
	fullName := Person{
		firstName: "john",
		lastName:  "aliu",
		age:       90.34,
	}
	// embedded struct
	secretAgent := Agent{
		Person: Person{
			firstName: "musa",
			lastName: "ALIU",
			age: 22,
		},
	}
     fmt.Println(secretAgent.lastName)
	fmt.Println("first name is", fullName.firstName, "and last name is", fullName.lastName, "and age is", fullName.age)
}
