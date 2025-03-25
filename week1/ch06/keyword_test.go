package ch06

import (
	"testing"
	"fmt"
	"reflect"
)

func TestKeyWord(t *testing.T){
	var a [3]int
	var b []int
	fmt.Println(reflect.TypeOf(a), "a=", a)
	fmt.Println(reflect.TypeOf(b), "b=", b, b==nil) 
	//切片的零值是nil 切片是引用类型，指向底层的数组，对引用的修改会影响到底层数组

	// append要赋值给新的变量
	b = append(b, 1,2,3)
	fmt.Println(reflect.TypeOf(b), "b=", b, b==nil)
}