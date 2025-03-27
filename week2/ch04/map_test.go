package main

import "fmt"

func main() {
    // 创建一个存储学生成绩的map
    studentScores := make(map[string]int)

    // 添加学生成绩
    studentScores["小明"] = 95
    studentScores["小红"] = 88
    studentScores["小李"] = 76

    // 打印所有学生成绩
    fmt.Println("学生成绩列表:")
    for name, score := range studentScores {
        fmt.Printf("%s: %d分\n", name, score)
    }

    // 修改小明的成绩
    studentScores["小明"] = 97
    fmt.Printf("\n修改后小明的成绩: %d分\n", studentScores["小明"])

    // 检查小王是否有成绩
    score, exists := studentScores['小王']
	if exists {
		fmt.Printf("\n小王的成绩:%d\n",score)
	}else {
		fmt.Println("小王的成绩不存在")
	}

    // 删除小李的成绩
    delete(studentScores, "小李")

    // 再次打印所有学生成绩
    fmt.Println("\n更新后的学生成绩列表:")
    for name, score := range studentScores {
        fmt.Printf("%s: %d分\n", name, score)
    }
}