package ch14

import (
	"fmt"
	"testing"
)

func panicFunc(){
	defer func ()  {
		if err := recover(); err != nil{
			println(err.(string))
		}
	}()
	panic("panic error!")
}

func TestPanic(t *testing.T){
	fmt.Println("before panic")
	panicFunc()
	fmt.Println("after panic")
}

func TestPanic2(t *testing.T){
	defer func ()  {
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()
	var ch chan int = make(chan int, 10) 
	close(ch)
	ch <- 1
}