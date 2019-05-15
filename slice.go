package main

import (
	"fmt"
)

func main() {
	slice()
	mapping()
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