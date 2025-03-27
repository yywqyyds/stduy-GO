package main

import "fmt"

func remove(slice []int) []int {
	a := make(map[int]bool)
	var result []int
	for _, v := range slice {
		if !a[v] {
			a[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}
	combinedSlice := append(slice1, slice2...)
	uniqueSlice := remove(combinedSlice)
	fmt.Print(uniqueSlice)
}
