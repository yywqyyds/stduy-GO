package ch07

import (
	"testing"

)

func TestMap(t *testing.T) {
	// 声明 + 初始化
	m1 := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	// 声明
	m2 := map[int]int{}
	m2[3] = 27
	t.Logf("len m2=%d", len(m2))

	// make 方法，可以设定 cap，但不能用 cap(m3)
	m3 := make(map[string]int, 10)
	t.Log(m3, len(m3))
}

func TestNonExist(t *testing.T){
	m3 := map[int]int{2:8}
	if val, ok := m3[3]; ok {
		t.Log("retrieved val =", val)
	}else {
		t.Log("key not found")
	}
}

func TestMapIteration(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9, 4: 16}
	for k,v := range m1{
		t.Log(k,"->",v)
	}
	delete(m1,4)
	for k,v := range m1{
		t.Log(k,"->",v)
	}
}
