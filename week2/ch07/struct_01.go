package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func NewPerson(name string, age int, email string) Person {
	return Person{Name: name, Age: age, Email: email}
}

func PrintPerson(p Person) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func main() {
	person := Person{
		Name:  "songhy",
		Age:   25,
		Email: "songhy@wust.edu.cn",
	}
	PrintPerson(person)
}
