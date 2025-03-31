package main

import "fmt"

func main() {
	var slice []int
	fmt.Println(slice, slice == nil)
	nil := [5]int{1, 2, 3, 4, 5}
	fmt.Println(nil)
	// fmt.Println(slice, slice == nil)
}
