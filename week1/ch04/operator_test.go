package ch04

import "testing"

func TestOperator(t *testing.T){
	//Go语言没有前置的 ++i,--i Python连后置也没有！
	var a = [3]int{1,2,3}
	var b = [3]int{1,2,4}
	t.Log(&a == &b)
	t.Log(a == b)
	b[2] = 3
	t.Log(&a == &b)
	t.Log(a == b)
}

func TestName(t *testing.T) {
	// 按位清零
	const (
		Readable = 1 << iota
		Writable
		Executable
	)

	a := 7 // 111
	t.Log(
		a&Readable == Readable,
		a&Writable == Writable,
		a&Executable == Executable,
		a&^Readable == Readable,
		a&^Writable == Writable,
		a&^Executable == Executable,
	)
}