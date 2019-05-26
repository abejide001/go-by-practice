package main

import (
	"fmt"
	"encoding/json"
	"os"
	"sort"
)
func main() {
	x := 10
	fmt.Println("this is where x is stored", &x) // using pointers
	fmt.Printf("%T\n", &x)
	b := &x
	fmt.Println("getting the value stored in x", *b)	
	// fmt.Println(odd(1,2,4,4,5,9,3))
	makeArray()
	sortArray()
	sliceArray()
	defer concat()
	slice()
	even()
	arr()
	mapping()
	obj()
	bar(22, 20)
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
	inc := incrementor()
	fmt.Println("incrementor was called", inc())
	fact := factorial(5)
	fmt.Println("factorial of 5 is", fact)
	factLoop := loop(5)
	fmt.Println("factorial of 5 is", factLoop)
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
	for i, v := range x {
		fmt.Println("the index is", i, "and the value is", v)
	}
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
	b, _ := json.Marshal(fullName)
	os.Stdout.Write(b)
	fullName.human() // human method
	getValue(fullName)
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
	getValue(secretAgent)
	fmt.Println("first name is", fullName.firstName, "and last name is", fullName.lastName, "and age is", fullName.age)
}

// method
func (s Person) human() {
	fmt.Println("alade is ", s.age)
}

// passing in a param
func bar(y int, z int) {
	fmt.Println(y, z)
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
		if array[i] % 2 == 0 {
			sum += array[i]
		}
	}
	fmt.Println("sum of even is ", sum)
}

// concat arrays
func concat() {
	arr1 := []int{1,2,3,5,6}
	arr2 := []int{1,2,2,3,3}
	conc := append(arr1, arr2...)
	fmt.Println("concatenating arrays ", conc)

	// delete from a slice
	arr1 = append(arr1[:4], arr1[1:]...) // this removes the item in the first index
	fmt.Println("returned array", arr1)
}

// slicing array 
func sliceArray() {
	x := []int{1,2,3,4,5,6,76}
	fmt.Println("slicng arrays", x[0:4])
}

// using make
func makeArray() {
	x := make([]int, 5, 5)
	fmt.Println("array with lenght 5", x)
	x = append(x, 20)
	fmt.Println(x)
	fmt.Println(cap(x))

}

// Interfaces(allow us to define behaviour)
type being interface {
	human()
}

func getValue(b being) {
	fmt.Println("i am using the method human", b)
}

// using callback functions
// func callback(f func(xi ...int) int, yi ...int) int { // passing in a func as an argument
//    var y []int
//    for _, v := range yi {
// 	   if v % 2 == 0 {
// 		   y = append(y, v)
// 	   }
//    }
//    return f(y...)
// }

// using closures
func incrementor() func() int {
  var x int
  return func() int {
	  x++
	  return x
  }
}

//using recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}

// recusrion with a loop
func loop(n int) int {
  total := 1
  for ; n > 0; n -- {
	  total *= n
  }
  return total
}

// sort an array 
func sortArray() {
	xi := []string{"voo", "adfsf", "bav"}
	// sort.Ints(xi)
	sort.Strings(xi)
	fmt.Println("sorted", xi)
}