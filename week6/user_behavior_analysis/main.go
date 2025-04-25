package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

// 用户状态结构体
type UserStat struct {
	Count     int
	Actions   map[string]int
	FirstSeen time.Time
	LastSeen  time.Time
}

// 分钟状态结构体
type MinuteStat struct {
	ActiveUsers map[string]bool
	TotalOps    int
}

func WirteUserStats(filename string, stats map[string]*UserStat) {
	//创建CSV文件
	file, _ := os.Create(filename)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	//插入内容
	writer.Write([]string{"User", "TotalOps", "FirstSeen", "LastSeen", "Login", "View", "Logout"})
	for user, s := range stats {
		writer.Write([]string{
			user,
			fmt.Sprintf("%d", s.Count),
			s.FirstSeen.Format("2006-01-02 15:04:05"),
			s.LastSeen.Format("2006-01-02 15:04:05"),
			fmt.Sprintf("%d", s.Actions["login"]),
			fmt.Sprintf("%d", s.Actions["view"]),
			fmt.Sprintf("%d", s.Actions["logout"]),
		})
	}
}

func WriteActionStats(filename string, stats map[string]int) {
	file, _ := os.Create(filename)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Action", "TotalCount"})
	for action, count := range stats {
		writer.Write([]string{action, fmt.Sprintf("%d", count)})
	}
}

func WriteMinuteStats(filename string, stats map[string]*MinuteStat) {
	file, _ := os.Create(filename)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Minute", "ActiveUsers", "TotalOps"})
	for minute, s := range stats {
		writer.Write([]string{
			minute,
			fmt.Sprintf("%d", len(s.ActiveUsers)),
			fmt.Sprintf("%d", s.TotalOps),
		})
	}
}

func main() {
	layout := "2006-01-02 15:04:05"
	userStats := make(map[string]*UserStat)
	actionStats := make(map[string]int)
	minuteStats := make(map[string]*MinuteStat)

	//打开日志文件
	file, err := os.Open("user_actions.log")
	if err != nil {
		fmt.Println("读取日志文件失败:", err)
		return
	}
	defer file.Close()

	//逐行读取文件
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")
		if len(fields) != 4 {
			continue
		}

		timestampStr, user, action := fields[0], fields[1], fields[2]
		timestamp, err := time.Parse(layout, timestampStr)
		if err != nil {
			fmt.Println("时间解析失败:", err)
			continue
		}
		minute := timestamp.Format("2006-01-02 15:04")

		// 用户统计
		if _, ok := userStats[user]; !ok {
			userStats[user] = &UserStat{
				Actions:   make(map[string]int),
				FirstSeen: timestamp,
				LastSeen:  timestamp,
			}
		}
		u := userStats[user]
		u.Count++
		u.Actions[action]++
		if timestamp.Before(u.FirstSeen) {
			u.FirstSeen = timestamp
		}
		if timestamp.After(u.LastSeen) {
			u.LastSeen = timestamp
		}

		// 行为统计
		actionStats[action]++

		// 分钟统计
		if _, ok := minuteStats[minute]; !ok {
			minuteStats[minute] = &MinuteStat{
				ActiveUsers: make(map[string]bool),
			}
		}
		m := minuteStats[minute]
		m.ActiveUsers[user] = true
		m.TotalOps++
	}

	//写入结果
	WirteUserStats("user_stats.csv", userStats)
	WriteActionStats("action_stats.csv", actionStats)
	WriteMinuteStats("minute_stats.csv", minuteStats)
}
