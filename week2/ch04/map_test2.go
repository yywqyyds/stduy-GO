package main

import "fmt"

func main() {
	// 创建一个存储学生成绩的map
	studentScores := make(map[string]int)

	// 添加学生成绩
	studentScores["小明"] = 94
	studentScores["小红"] = 88
	studentScores["小李"] = 76
	studentScores["小王"] = 100
	studentScores["小宋"] = 99
	// 打印所有学生成绩
	fmt.Println("学生成绩列表:")
	for i := 0; i < 100; i++ {
		for name, _ := range studentScores {
			fmt.Printf("%s", name)
		}
		fmt.Println()
	}
}
