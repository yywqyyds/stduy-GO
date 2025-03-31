package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type Q struct {
	b byte
	i int64
	u uint16
}

type S struct {
	b byte
	u uint16
	i int64
}

func TestStructSize(t *testing.T) {
	var q Q
	fmt.Println(unsafe.Sizeof(q))
	var s S
	fmt.Println(unsafe.Sizeof(s))
}