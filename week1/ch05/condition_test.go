package ch05

import (
	"testing"
	"fmt"
)

func TestIf(t *testing.T){
	a := 10
	if a > 9 {
		fmt.Printf("%d是一个大于9的数",a)
	}
}

func TestSwitch(t *testing.T){
	// switch 不限于常量整形，可以是字符串
	a:= 2
	switch a{
	case 1: 
		fmt.Print("输出数字1")
	case 2:
		fmt.Print("输出数字2")
	}

}

func TestSwitch2(t *testing.T){
	isMatch := func (i int) bool{
		switch i{
			case 1:
				fallthrough //穿透keyword，继续执行下一个case语句，而不管下一个case条件是否满足
			case 2:
				return true
		}
		return false
	}
	fmt.Print(isMatch(1))
	fmt.Print(isMatch(2))
}