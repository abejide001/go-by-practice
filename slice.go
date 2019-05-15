package main

import (
	"fmt"
)

func main() {
	slice()
	mapping()
}
func slice() {
	// y := []string{"money", "cash", "adsd", "dsfsf", "fdgggg"}
	// x := []string{"money", "cash", "adsd", "dsfsf", "fdgggg", "ghost"}
	// rem := [][]string{x, y}
	// fmt.Println(rem)
	x := []int{1,3,24,5,6,76}
	y := append(x[1:3])
	fmt.Println(y)
}
type Person struct {
	first string
	last string
}

func mapping() {
	// x := map[string]int{   
	// 	"John": 90,
	// }
	// x["Mumum"] = 20
	// delete(x, "Mumum")
	// fmt.Println(x)
	p1 := Person{
		first: "Boss",
		last: "fsfdg",
	}
   fmt.Println(p1.first)
}