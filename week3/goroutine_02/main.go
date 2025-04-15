package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func countWord(filePath string, word string, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("无法打开文件%s:%v\n", filePath, err)
		ch <- 0
		return
	}
	defer file.Close()
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, w := range words {
			if strings.EqualFold(w, word) {
				count++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("读取文件%s出错:%v\n", filePath, err)
	}
	ch <- count
}

func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	targetWord := "hello"

	var wg sync.WaitGroup
	ch := make(chan int, len(files))

	for _, file := range files {
		wg.Add(1)
		go countWord(file, targetWord, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	total := 0
	for c := range ch {
		total += c
	}

	fmt.Printf("单词%s在所有文件中总共出现了%d次", targetWord, total)
}
