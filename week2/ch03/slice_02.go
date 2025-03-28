package main

import "fmt"

func main() {
	slice := make([]int, 10)
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := slice[2:7]
	slice1 = append(slice1, 11, 12, 13)
	//删除第5个元素
	slice1 = append(slice1[:4], slice1[5:]...)
	//将切片中所有元素乘以2
	for i := range slice1 {
		slice1[i] *= 2
	}
	fmt.Println(slice1, cap(slice1))
}
