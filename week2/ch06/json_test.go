package ch06

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 结构体转换为json
func TestJson(t *testing.T) {
	type Person struct {
		Name    string
		Age     int
		Hobbies []string
	}
	person := Person{
		Name:    "John",
		Age:     25,
		Hobbies: []string{"reading", "writing", "coding"},
	}
	//将Person结构体转换为json格式的字节切片
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(string(jsonData))
}

// json转换为结构体
func TestJson2(t *testing.T) {
	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Hobbies []string `json:"hobbies"`
	}
	jsonData := `{"name":"John Doe","age":30,"hobbies":["reading","writing","coding"]}`
	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 输出Person结构体的字段值
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Hobbies:", person.Hobbies)
}
