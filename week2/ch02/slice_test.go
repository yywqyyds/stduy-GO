package ch02

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var arr = [5]int{1, 2, 3, 4, 5}
	arr[0] = 2
	var slice = arr[2:4]
	slice[0] = 2
	fmt.Print(arr)
}
