package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("欢迎来到猜数字游戏")
	fmt.Println("下面是游戏规则：\n1.计算机将在1到100之间随机选择一个数字。\n2.您可以选择难度级别（简单,中等，困难），不同难度对应不同的猜测机会。\n3.请输入您的猜测。")

	// 记录文件
	file, err := os.OpenFile("game.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("无法创建记录文件：", err)
		return
	}
	defer file.Close()

	if stat, _ := file.Stat(); stat.Size() == 0 {
		file.WriteString("游戏记录\n")
		file.WriteString("时间 | 难度 | 猜测次数 | 结果 | 用时 | 正确答案\n")
	}

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	for {
		var number = rand.Intn(100) + 1
		difflevel := map[int]int{
			1: 10,
			2: 5,
			3: 3,
		}
		var difficulty int

		for {
			fmt.Println("请选择难度级别（简单/中等/困难):\n1.简单(10次机会)\n2.中等(5次机会)\n3.困难(3次机会)")
			if _, err := fmt.Scan(&difficulty); err != nil {
				fmt.Println("输入无效，请输入1、2 或 3")
				continue
			}
			if difficulty < 1 || difficulty > 3 {
				fmt.Println("输入错误，请输入 1、2 或 3")
				continue
			}
			break
		}

		guessTimes := difflevel[difficulty]
		var guessNumber int
		startTime := time.Now()
		flag := false
		var attempt int
		var record string

		fmt.Printf("您选择了%d级难度,您有%d次猜测机会。\n", difficulty, guessTimes)

		for i := 1; i <= guessTimes; i++ {
			fmt.Printf("第%d次猜测,请输入您的数字(1-100):\n", i)
			_, err := fmt.Scan(&guessNumber)
			if err != nil || guessNumber < 1 || guessNumber > 100 {
				fmt.Println("请输入1-100之间的数字")
				i--
				continue
			}

			attempt = i

			if guessNumber > number {
				fmt.Println("您猜的数字大了。")
			} else if guessNumber < number {
				fmt.Println("您猜的数字小了。")
			} else {
				guessNumberTime := time.Since(startTime)
				fmt.Printf("恭喜您猜对了！您在第%d次猜测中成功。用时%.1f秒。\n", i, guessNumberTime.Seconds())
				flag = true
				record = fmt.Sprintf("[%s] | 难度%d | %d次 | 成功 | %.1f秒 | 答案:%d\n",
					time.Now().Format("2006-01-02 15:04:05"), difficulty, i, guessNumberTime.Seconds(), number)
				break
			}
		}

		if !flag {
			guessNumberTime := time.Since(startTime)
			fmt.Printf("很可惜您没有在规定的次数内猜对数字,用时%.1f秒,正确答案是%d,游戏结束,您可以退出或者重新开始。\n", guessNumberTime.Seconds(), number)
			record = fmt.Sprintf("[%s] | 难度%d | %d次 | 失败 | %.1f秒 | 正确答案:%d\n",
				time.Now().Format("2006-01-02 15:04:05"), difficulty, attempt, guessNumberTime.Seconds(), number)
		}

		if _, err := file.WriteString(record); err != nil {
			fmt.Println("保存记录失败:", err)
		}

		var playAgain string
		fmt.Println("是否再来一次？(y/n)")
		fmt.Scan(&playAgain)
		if playAgain == "n" {
			fmt.Println("游戏结束,感谢您的游玩！您的游戏记录已保存")
			break
		}
	}
}
