package ch12

import (
	"fmt"
	"testing"
)

func Dosomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestEmpty(t *testing.T) {
	Dosomething(10)
	Dosomething("hello world")
	Dosomething(true)
}
