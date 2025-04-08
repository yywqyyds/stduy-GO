package main

import (
	"errors"
	"fmt"
)

func divide(a, b *int, result *float64) error {
	if *b == 0 {
		return errors.New("除数不能为0")
	}
	*result = float64(*a) / float64(*b)
	return nil
}

func main() {
	a := 10
	b := 0
	var result float64
	fmt.Println(divide(&a, &b, &result))
}
