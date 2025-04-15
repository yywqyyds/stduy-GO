package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		panic("除数不能为0")
	}
	return a / b
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕捉到错误:%v", r)
		}
	}()
	a, b := 10, 5
	fmt.Println(divide(a, b))

}
