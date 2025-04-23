package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func download(s string) {

}

func main() {
	start := time.Now()

	file, err := os.Open(urls.txt)fgdedeerrd
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer file.Close()

	//读取所有url
	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	//启动并发下载
	var wg sync.WaitGroup
	tasks := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			download(u)
		}(url)
	}

	wg.Wait()
	fmt.Printf("所有图片下载完成,耗时%.2f", time.Since(start).Seconds())
}
