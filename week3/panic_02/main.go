package main

import "fmt"

func accessArray(arr []int, index int) int {
	if index < 0 || index > len(arr) {
		panic("索引超过数组范围")
	}
	return arr[index]
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕捉到错误:", r)
		}
	}()
	arr := [5]int{1, 2, 3, 4, 5}
	index := 6
	value := accessArray(arr[:], index)
	fmt.Printf("索引%d对应的值为:%d", index, value)
}
