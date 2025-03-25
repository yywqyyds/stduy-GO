package ch05

import (
	"testing"
	"fmt"

)
func TestFor(t *testing.T){
	for i:=0; i<10; i++{
		fmt.Println(i)
	}
}

func TestRange(t *testing.T){
	s := "hello Go"
	a := [3]int{1,2,3}
	for _,i := range s{
		fmt.Println(i)
	}
	for _,i := range a{
		fmt.Println(i)
	}
}