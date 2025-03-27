package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := slice[2:7]
	slice1 = append(slice1, 100)
	sum := 0
	for i, _ := range slice1 {
		sum += slice1[i]
	}
	fmt.Println(sum)
}
