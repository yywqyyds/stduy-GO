package main

import (
	"fmt"
	"sync"
)

func recv(ch chan int,wg *sync.WaitGroup) {
	defer wg.Done()
	for c := range ch {
		fmt.Printf("接收者:接收到数据%d\n", c)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("生产者：生产数字%d\n", i)
			ch <- i
		}
		close(ch)
	}()
	go recv(ch,&wg)
	wg.Wait()
}
