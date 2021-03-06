package main

import (
	"encoding/json"
	f "fmt"
	"os"
	"runtime"
	"sort"
	"sync"
	s "strings"
)
var wg sync.WaitGroup
func main() {
	value := 42
	point(&value)
	// channel()
	f.Println("number of goroutine is", runtime.NumGoroutine())
	f.Println("NUMBER OF CPU", runtime.NumCPU())
	f.Println("cgo call", runtime.NumCgoCall())
	wg.Add(1) // takes an integer
	go fooo() // launc the function into a goroutine
	barr()
	f.Println("number of goroutine is", runtime.NumGoroutine())
	wg.Wait() // wait for the execution
	unMarsh()
	anim()
	x := "wale"
	f.Println("this is where x is stored", &x) // using pointers
	f.Printf("%T\n", &x)
	b := &x // reassigning to variable b
	f.Println("getting the value stored in x", *b)
	// f.Println(odd(1,2,4,4,5,9,3))
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
	f.Println(anon.color)
	s1 := woo("dsfsf")
	f.Println(s1)
	// anon function
	func() {
		f.Println("anon")
	}()
	inc := incrementor()
	f.Println("incrementor was called", inc())
	fact := factorial(5)
	f.Println("factorial of 5 is", fact)
	factLoop := loop(5)
	f.Println("factorial of 5 is", factLoop)
}

// using slice
func slice() {
	y := []string{"money", "cash", "adsd", "dsfsf", "fdgggg"}
	x := []string{"money", "cash", "adsd", "dsfsf", "fdgggg", "ghost"}
	rem := [][]string{x, y}
	f.Println(rem)
}

// mapping
func mapping() {
	x := map[string]int{
		"John": 90,
	}
	x["Mumum"] = 20
	for i, v := range x {
		f.Println("the index is", i, "and the value is", v)
	}
	f.Println(x)
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
	f.Println(fullName)
	// embedded struct
	secretAgent := Agent{
		Person: Person{
			firstName: "musa",
			lastName:  "ALIU",
			age:       22,
		},
	}
	f.Println(secretAgent.lastName)
	getValue(secretAgent)
	f.Println("first name is", fullName.firstName, "and last name is", fullName.lastName, "and age is", fullName.age)
}

// method
func (s Person) human() {
	f.Println("alade is ", s.age)
}

// passing in a param
func bar(y int, z int) {
	f.Println(y, z)
}

// using a return
func woo(y string) string {
	return f.Sprint("Hello ", y)
}

// getting the sum of digits in an  array
func arr() {
	sum := 0
	array := []int{1, 2, 3, 4}
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	f.Println("sum is stored here", &sum)
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
	f.Println("sum of even is ", sum)
}

// concat arrays
func concat() {
	arr1 := []int{1, 2, 3, 5, 6}
	arr2 := []int{1, 2, 2, 3, 3}
	conc := append(arr1, arr2...)
	f.Println("concatenating arrays ", conc)

	// delete from a slice
	arr1 = append(arr1[:4], arr1[1:]...) // this removes the item in the first index
	f.Println("returned array", arr1)
}

// slicing array
func sliceArray() {
	x := []int{1, 2, 3, 4, 5, 6, 76}
	f.Println("slicng arrays", x[0:4])
}

// using make
func makeArray() {
	x := make([]int, 5, 5)
	f.Println("array with lenght 5", x)
	x = append(x, 20)
	f.Println(x)
	f.Println(cap(x))

}

// Interfaces(allow us to define behaviour)
type being interface {
	human()
}

func getValue(b being) {
	f.Println("i am using the method human", b)
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
	return n * factorial(n-1)
}

// recusrion with a loop
func loop(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

// sort an array
func sortArray() {
	xi := []string{"voo", "adfsf", "bav"}
	// sort.Ints(xi)
	sort.Strings(xi)
	f.Println("sorted", xi)
}

// trying byte
func typeByte() {
	var x byte = 22
	f.Println(x)
}

// convert strings to lower case
func convert() {
	x := "abcde"
	x = s.ToLower(x)
	f.Println(x)
}

// Animal struct(using marshal)
type Animal struct {
	Color  string `json:"Color"`
	Height int    `json:"Height"`
}

// json marshalling
func anim() {
	a := Animal{
		Color:  "black",
		Height: 5,
	}
	x := []Animal{a} // slice of type animal
	b, err := json.Marshal(x)
	if err != nil {
		f.Println(err)
	}
	f.Println(string(b))
}

// json unmarshalling
func unMarsh() {
	s := `[{"Color":"black","Height":5}]`
	bs := []byte(s)
	var a []Animal
	err := json.Unmarshal(bs, &a)
	if err != nil {
		f.Println(err)
	}
	f.Println(a)
}

// understanding concurrency
func fooo() {
	for i := 0; i < 10; i++ {
		f.Println("foo", i)
	}
}

func barr() {
	for i := 0; i < 10; i++ {
		f.Println("bar", i)
	}
	wg.Done()
}

// using channels
// func channel() {
// 	c := make(chan string, 3) // buffer channel
// 	c <- "4ww2"
// 	c <- "43"
// 	c <- "44"
// 	f.Println(<-c)
// 	f.Println(<-c) // pulling of a value from a channel
// 	f.Println(<-c)

// }

func point(n *int) {
  f.Println("should return a an address", n)
}