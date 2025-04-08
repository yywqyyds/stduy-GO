package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	Completed   bool      `json:"completed"`
	CreatedTime time.Time `json:"created_time"`
}

type TaskManager struct {
	Tasks []Task `json:"tasks"` // 人物列表
	file  string
}

func (tm *TaskManager) loadTasks() error {
	if _, err := os.Stat(tm.file); os.IsNotExist(err) {
		tm.Tasks = []Task{}
		return nil
	}
	//读取文件内容
	data, err := ioutil.ReadFile(tm.file)
	if err != nil {
		return err
	}
	//反序列化json
	return json.Unmarshal(data, &tm.Tasks)
}

func (tm *TaskManager) SaveTasks() error {
	//序列化为带缩进的json
	data, err := json.Marshal(tm.Tasks, "", " ")
	if err != nil {
		return nil
	}
	//写入文件
	return ioutil.WriteFile(tm.file, data, 0644)
}

func (tm *TaskManager) AddTask(content string) {
	newId := 1
	if len(tm.Tasks) > 0 {
		newId = tm.Tasks[len(tm.Tasks)-1].Id + 1
	}

	//创建新任务
	task := Task{
		Id:          newId,
		Content:     content,
		Completed:   false,
		CreatedTime: time.Now(),
	}
	tm.Tasks = append(tm.Tasks, task)
	err := tm.SaveTasks()
	if err != nil {
		fmt.Printf("保存任务失败：%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("已添加任务#%d: %s\n", task.Id, task.Content)
}

func (tm *TaskManager) DoneTasks(Id int){
	for i := range tm.Tasks{
		if tm.Tasks[i].Id == Id{
			//标记为完成并保存
			tm.Tasks[i].Completed = true
			err := tm.SaveTasks()
			if err != nil{
				fmt.Printf("保存任务失败:%v\n",err)
				os.Exit(1)
			}
			fmt.Printf("已完成任务#%d:%s\n",Id,tm.Tasks[i].Content)
		}
	}
	fmt.Printf("错误:找不到Id为%d的任务\n",Id)
	os.Exit(1)
}

func (tm *TaskManager) DeleteTasks(Id int){
	//遍历查找任务
	for i, task := range tm.Tasks{
		if task.Id == Id{
			//删除元素
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			err := tm.SaveTasks()
			if err != nil{
				fmt.Printf("保存任务失败:%v\n",err)
				os.Exit(1)
			}
			fmt.Printf("已删除任务 #%d: %s\n", Id, task.Content)
			return
		}
	}
	//未找到任务
	fmt.Printf("错误：找不到Id为%d的任务\n",Id)
	os.Exit(1)
}

func (tm *TaskManager) ListTasks(){
	//显示所有未完成任务
	unCompletedTasks := []Task{}
	for _,task := range tm.Tasks{
		if task.Completed == false{
			unCompletedTasks = append(unCompletedTasks, task)
		}
	}
	if len(unCompletedTasks) == 0{
		fmt.Println("没有未完成的任务")
		return
	}
	fmt.Println("未完成任务列表")
	for _, task := range unCompletedTasks{
		fmt.Printf("#%d %s(创建于:%s)\n",
		task.Id,
		task.Content,
		task.CreatedTime.Format("2006-01-02 15:04"))
	}
}

func main() {
	tm := TaskManager{file: "task.json"}
	err := tm.loadTasks()
	if err != nil {
		fmt.Printf("任务加载失败：", err)
		os.Exit(1)
	}
	//设置子命令
	addCmd := flag.NewFlagSet("add",flag.ExitOnError)
	addContent := addCmd.String("content", "", "任务内容")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	list := listCmd.Bool("")

	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	completeId := completeCmd.Int("Id",0,"要完成的任务Id")

	deleteCmd := flag.NewFlagSet("delete",flag.ExitOnError)
	deleteId := deleteCmd.Int("Id",0,"要删除的任务Id")

	if len(os.Args) < 2{
		os.Exit(1)
	}

	//解析子命令
	switch os.Args[1]{
	case "add":
		addCmd.Parse(os.Args[2:])
	}
}
