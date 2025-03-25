package ch03

import "testing"

func TestType(t *testing.T){
	// Go对于类型转换很严格，只支持显示转换
	var a int32 = 1
	var b = int64(a)
	t.Log(a,b)
}