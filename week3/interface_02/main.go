package main

import "fmt"

type Animal interface {
	Speak() string
	Move() string
	Name() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}
type Bird struct {
	name string
}

func (d Dog) Speak() string {
	return "汪汪汪!"
}

func (c Cat) Speak() string {
	return "喵喵喵~"
}

func (b Bird) Speak() string {
	return "咕咕咕?"
}

func (d Dog) Move() string {
	return "四脚跑"
}

func (c Cat) Move() string {
	return "四脚跑"
}

func (b Bird) Move() string {
	return "翅膀飞"
}

func (d Dog) Name() string {
	return d.name
}

func (c Cat) Name() string {
	return c.name
}

func (b Bird) Name() string {
	return b.name
}

func main() {
	animals := []Animal{
		Dog{name: "旺财"},
		Cat{name: "果冻"},
		Bird{name: "飞飞"},
	}
	for _, animal := range animals {
		fmt.Printf("%s说: %s, %s的移动方式: %s\n", animal.Name(), animal.Speak(), animal.Name(), animal.Move())
	}
}
