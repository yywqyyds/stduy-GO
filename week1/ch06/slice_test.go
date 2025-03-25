package ch06

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T){
	var s0 []int //空切片
	var s1 = make([]int, 5,10) // make创建的切片全是0
	fmt.Println("切片s0:",s0)
	fmt.Println("切片s1:",s1)
	s0 = append(s0, 1,2,3)
	fmt.Println("添加后的切片s0:",s0)
	s1 = append(s1, 1,2,3,4,5)
	fmt.Println("添加后的切片s1:",s1)
}

func TestSlice2(t *testing.T){
	month := []string{
		"Jan", "Feb", "Mar", "Apr", "May",
		"Jun", "Jul", "Aug", "Sep", "Oct", "Nov",
		"Dec",
	}
	spring := month[0:3]
	summer := month[3:6]
	fmt.Println(spring,summer)
}