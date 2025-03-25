package ch09

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	// string不是引用类型，空值为""
	var s string = "songhaoyuan"
	fmt.Print(s, len(s))
	chars := []rune(s) //将字符串转换为rune切片  rune本质上是int32
	fmt.Print(chars)
}

func TestString2(t *testing.T) {
	str := "hello Go"
	// cannot assign to str[0] (value of type byte)因为字符串是不可变的
	//str[0] = 'x'
	fmt.Println(str)
}

func TestString3(t *testing.T) {
	str := "hello Go"
	for _, c := range str {
		fmt.Println(c)
	}
}
