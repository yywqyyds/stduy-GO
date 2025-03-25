package ch06

import (
	"testing"
)

func TestArray(t *testing.T){
	var a [3]int
	t.Log(a)
	b := [...]int{1,2,3,4,5,6,7,8}
	for _, val := range b{
		t.Log(val)
	}
	// 数组截取与python类似，但是不支持步进，也不支持-1等倒数元素
	t.Log(b[1:])
}