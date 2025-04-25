package main

import (
	"testing"
)

func TestChan(t *testing.T) {
	ch := make(chan int, 2)
	a := 0
	ch <- 2
	close(ch)
	a = <-ch
	println(a)

}
