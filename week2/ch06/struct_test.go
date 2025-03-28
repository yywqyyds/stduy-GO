package ch06

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
	Sex  string
	City string
}

// 同样类型的可以写到一行
type Person1 struct {
	Name, Sex, City string
	Age             int
}

func (p *Person) Agejiayi() {
	p.Age += 1
}

func TestStruct(t *testing.T) {
	//结构体实例化
	// var p1 Person
	p2 := Person{
		Name: "songhy",
		Age:  25,
		Sex:  "男",
		City: "wuhan",
	}
	// p3 := new(Person)
	//方法定义
	// func (p Person) Introduce(){
	// 	fmt.Printf("大家好我叫%s,今年%d岁,住在%s.\n",p.Name,p.Age,p.City)
	// }
	p2.Agejiayi()
	fmt.Println(p2.Age)
}
