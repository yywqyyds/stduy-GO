package main

import "fmt"

func main() {
	arr := "dsnajadasn45612456"
	result := make(map[rune]int)
	maxCharTimes := 0
	var maxChar rune
	for _, char := range arr {
		times, exists := result[char]
		if exists {
			result[char] = times + 1
		} else {
			result[char] = 1
		}
		if result[char] > maxCharTimes {
			maxCharTimes = result[char]
			maxChar = char
		}
	}
	fmt.Println(result)
	fmt.Printf("出现次数最多的字符是%c,次数为%d", maxChar, maxCharTimes)
}
