package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./data/questions.db")
	if err != nil {
		log.Fatalf("打开数据库失败: %v", err)
	}

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS questions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        model TEXT,
        language TEXT,
        type INTEGER,
        keyword TEXT,
        difficulty TEXT,
        question TEXT,
        options TEXT,
        answer TEXT,
        explanation TEXT,
        ai_start_time TEXT,
        ai_end_time TEXT,
        ai_cost_time INTEGER
    );`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("创建表失败: %v", err)
	}
}
