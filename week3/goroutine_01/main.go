package main

import (
	"fmt"
	"sync"
)

func getSums(nums []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	sum := 0
	for _, num := range nums {
		sum += num * num
	}
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	parts := 3
	partSize := (len(nums) + parts - 1) / parts

	var wg sync.WaitGroup
	ch := make(chan int, parts)
	for i := 0; i < parts; i++ {
		start := i * partSize
		end := start + partSize
		if end > len(nums) {
			end = len(nums)
		}
		wg.Add(1)
		go getSums(nums[start:end], &wg, ch)
	}

	wg.Wait()
	close(ch)

	total := 0
	for c := range ch {
		total += c
	}

	fmt.Printf("整个切片的平方和为:%d\n", total)
}
