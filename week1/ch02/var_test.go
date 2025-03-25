package ch02

import (
	"fmt"
	"testing"
)

func TestVar(t *testing.T){
	var a int = 1
	var b = 2
	c := 3
	var str string = "adawdf"
	var arr = [...]int{1,2,3}
	fmt.Println(a,b,c,str,arr)
}