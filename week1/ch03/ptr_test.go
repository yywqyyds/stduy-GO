package ch03

import (
	"testing"
)

func TestPointer(t *testing.T) {
	// 不支持指针操作
	var a = 1
	var aPtr = &a
	t.Log(a, aPtr)

}
