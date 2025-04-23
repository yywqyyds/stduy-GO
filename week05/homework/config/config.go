package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

var (
	ModelName string
	ApiKey string
)

//加载.env文件
func LoadConfig(path string)error {
	err := godotenv.Load(path)
	if err != nil{
		return err
	}
	ModelName = os.Getenv("MODEL")
	ApiKey = os.Getenv("API_KEY")
	if ModelName == "" || ApiKey == ""{
		log.Panicln("MODEL或者API_KEY未设置")
	}
	return nil
}