package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
	"time"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 查找指定范围内的素数
func findPrimes(start, end int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if isPrime(i) {
			out <- i
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("用法：go run main.go start end")
		return
	}

	startIn, err1 := strconv.Atoi(os.Args[1])
	endIn, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || startIn > endIn {
		fmt.Println("参数错误：请输入两个整数，且 start <= end")
		return
	}

	runPrimeFinder(startIn, endIn)
}

func runPrimeFinder(startIn, endIn int) {
	startTime := time.Now()

	threads := 4
	total := endIn - startIn + 1
	step := total / threads
	if step == 0 {
		step = 1
	}

	resultChan := make(chan int, 1000)
	var wg sync.WaitGroup

	for i := 0; i < threads; i++ {
		start := startIn + i*step
		end := startIn + (i+1)*step - 1
		if start > endIn {
			break
		}
		if end > endIn {
			end = endIn
		}
		wg.Add(1)
		go findPrimes(start, end, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	primes := []int{}
	for p := range resultChan {
		primes = append(primes, p)
	}

	elapsed := time.Since(startTime).Milliseconds()

	if len(primes) == 0 {
		fmt.Printf("未找到素数。用时：%d 毫秒\n", elapsed)
		return
	}

	fileName := fmt.Sprintf("prime_results_%d_%d.txt", startIn, endIn)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer file.Close()

	for _, p := range primes {
		fmt.Fprintln(file, p)
	}

	fmt.Printf("一共找到了 %d 个素数，耗时 %d 毫秒，已写入 %s\n", len(primes), elapsed, fileName)
}
