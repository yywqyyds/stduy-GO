package ch01

import (
	"fmt"
	"testing"
)

func TestPlace(t *testing.T) {
	a := 10
	p := &a
	fmt.Print(p)
}

func TestPlace2(t *testing.T) {
	aint := 1
	astring := "hello Go"
	var ptr *int
	ptr = &aint
	fmt.Printf("%T\n", ptr) //%T数据类型
	fmt.Printf("%T\n", &astring)
	fmt.Println("普通变量存储的是：", aint)
	fmt.Println("普通变量存储的是：", *ptr)
	fmt.Println("指针变量存储的是：", &aint)
	fmt.Println("指针变量存储的是：", ptr)
}
