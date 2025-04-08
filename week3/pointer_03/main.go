package main

import "fmt"

func doubleValues(p *[5]int) {
	for i := 0; i < len(p); i++ {
		p[i] *= 2
	}
}

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	doubleValues(&nums)
	fmt.Println(nums)
}
