package ch02

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	// 声明并初始化一周的温度数组（单位：摄氏度）
	weeklyTemps := [7]float64{28.5, 30.2, 26.8, 27.5, 31.0, 29.8, 27.2}

	// 打印每天温度
	fmt.Println("一周温度记录：")
	days := [7]string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
	for i, temp := range weeklyTemps {
		fmt.Printf("%s: %.1f°C\n", days[i], temp)
	}

	// 计算平均温度
	sum := 0.0
	for _, temp := range weeklyTemps {
		sum += temp
	}
	average := sum / float64(len(weeklyTemps))
	fmt.Printf("\n平均温度: %.1f°C\n", average)

	// 找出最高和最低温度
	maxTemp := weeklyTemps[0]
	minTemp := weeklyTemps[0]
	maxDay := 0
	minDay := 0

	for i, temp := range weeklyTemps {
		if temp > maxTemp {
			maxTemp = temp
			maxDay = i
		}
		if temp < minTemp {
			minTemp = temp
			minDay = i
		}
	}

	fmt.Printf("最高温度: %.1f°C (%s)\n", maxTemp, days[maxDay])
	fmt.Printf("最低温度: %.1f°C (%s)\n", minTemp, days[minDay])

	// 计算温差
	tempDiff := maxTemp - minTemp
	fmt.Printf("一周温差: %.1f°C\n", tempDiff)
}
