package ch03

import (
	"fmt"
	"testing"
)

func mytest(ptr *int) {
	fmt.Println(*ptr)
}

func modify(sls []int) {
	sls[0] = 90
}

func TestPoint(t *testing.T) {
	// 创建指针
	aint := 100
	ptr := new(string)
	ptr1 := &aint
	var ptr2 *int
	ptr2 = &aint
	fmt.Print(ptr, *ptr1, ptr2)
	fmt.Printf("指针的类型是：%T\n", &ptr)
	mytest(&aint)
}

func TestPoint1(t *testing.T) {
	a := 25
	var b *int

	if b == nil {
		fmt.Println(b)
		b = &a
		fmt.Println(*b)
	}
}

func TestPoint2(t *testing.T) {
	arr := [3]int{89, 90, 91}
	modify(arr[:])
	fmt.Println(arr)
}
