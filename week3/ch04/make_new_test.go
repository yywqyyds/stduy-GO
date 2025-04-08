package ch04

import (
	"fmt"
	"testing"
)

func TestMake(t *testing.T) {
	//初始化切片
	slice := make([]int, 3)
	fmt.Printf("slice:%v\n", slice)
	//初始化map
	m := make(map[string]int)
	m["key"] = 1
	fmt.Printf("map:%v\n", m)
	//new示例
	numPtr := new(int)
	fmt.Printf("Pointer to int: %v, Value: %d\n", numPtr, *numPtr)
	type MyStrucr struct {
		a int
		b string
	}

	structPtr := new(MyStrucr)
	fmt.Printf("Pointer to struct: %v, a: %d, b: %s\n", structPtr, structPtr.a, structPtr.b)
}
