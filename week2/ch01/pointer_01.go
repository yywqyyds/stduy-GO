package main

import "fmt"

func Swap(a, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

func main() {
	num1 := 5
	num2 := 10
	fmt.Printf("交换前: num1 = %d,num2 = %d\n", num1, num2)
	Swap(&num1, &num2)
	fmt.Printf("交换后: num1 = %d,num2 = %d\n", num1, num2)
}
