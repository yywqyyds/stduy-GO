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
		fmt.Println("输入：go run main.go start end")
		return
	}
	startIn, err1 := strconv.Atoi(os.Args[1])
	endIn, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || startIn > endIn {
		fmt.Println("输入参数错误,请输入两个整数且start<=end")
		return
	}

	startTime := time.Now() //记录开始时间

	threads := 4
	step := (endIn - startIn + 1) / threads

	//创建通道用于传递素数
	resultChan := make(chan int, 1000)
	var wg sync.WaitGroup

	//启动4个协程并发查找素数
	for i := 0; i < threads; i++ {
		start := startIn + i*step + 1
		end := startIn + (i+1)*step
		if i == threads-1 {
			end = endIn
		}
		wg.Add(1)
		go findPrimes(start, end, resultChan, &wg)
	}

	//启动一个协程关闭通道（等所有任务完成）
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	count := 0
	for p := range resultChan {
		fmt.Fprintln(file, p)
		count++
	}

	//写入文件
	fileName := fmt.Sprintf("prime_%d_%d.txt", startIn, endIn)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer file.Close()

	elapased := time.Since(startTime).Seconds()
	fmt.Printf("一共找到了%d个素数,用时：%.2f\n", count, elapased)
}
