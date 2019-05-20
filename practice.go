package main

import (
	"fmt"
)

func main() {
	concat()
	slice()
	even()
	arr()
	mapping()
	obj()
	bar(22)
	// anonymous struct
	anon := struct {
		color string
	}{
		color: "black",
	}
	fmt.Println(anon.color)
	s1 := woo("dsfsf")
	fmt.Println(s1)
	// anon function
	func() {
		fmt.Println("anon")
	}()
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
	fullName.human()
	fmt.Println(fullName)
	// embedded struct
	secretAgent := Agent{
		Person: Person{
			firstName: "musa",
			lastName:  "ALIU",
			age:       22,
		},
	}
	fmt.Println(secretAgent.lastName)
	fmt.Println("first name is", fullName.firstName, "and last name is", fullName.lastName, "and age is", fullName.age)
}

// method
func (s Person) human() {
	fmt.Println("alade is ", s.age)
}

// passing in a param
func bar(y int) {
	fmt.Println(y)
}

// using a return
func woo(y string) string {
	return fmt.Sprint("Hello ", y)
}

// getting the sum of digits in an  array
func arr() {
	sum := 0
	array := []int{1, 2, 3, 4}
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	fmt.Println(sum)
}

// getting the sum of even digits in an  array
func even() {
	sum := 0
	array := []int{1, 2, 3, 4}

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			sum += array[i]
		}
	}
	fmt.Println("sum of even is ", sum)
}

// concat arrays
func concat() {
	arr1 := []int{1,2,2,3,3}
	arr2 := []int{1,2,2,3,3}
	conc := append(arr1, arr2...)
	fmt.Println("concatenating arrays ", conc)
}