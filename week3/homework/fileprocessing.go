package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

type Translation struct {
	Type        string `json:"type"`
	Translation string `json:"translation"`
}

type Phrase struct {
	Phrase      string `json:"phrase"`
	Translation string `json:"translation"`
}

type Word struct {
	Word         string        `json:"word"`
	Translations []Translation `json:"translations"`
	Phrases      []Phrase      `json:"phrases"`
}

func createTables(db *sql.DB) {
	//单词表
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS words (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		word TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal("创建单词表失败:", err)
	}

	// 词性表
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS translations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		word_id INTEGER,
		type TEXT,
		translation TEXT,
		FOREIGN KEY(word_id) REFERENCES words(id)
	);`)
	if err != nil {
		log.Fatal("创建词性表失败:", err)
	}

	// 短语表
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS phrases (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		word_id INTEGER,
		phrase TEXT,
		translation TEXT,
		FOREIGN KEY(word_id) REFERENCES words(id)
	);`)
	if err != nil {
		log.Fatal("创建短语表失败:", err)
	}
}

func insertData(db *sql.DB, words []Word) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("开启事务失败:", err)
	}
	//预编译语句
	wordSt, _ := tx.Prepare(`INSERT OR IGNORE INTO words(word) VALUES(?)`)
	translationSt, _ := tx.Prepare(`INSERT INTO translations(word_id, type, translation) VALUES (?, ?, ?)`)
	phraseSt, _ := tx.Prepare(`INSERT INTO phrases(word_id, phrase, translation) VALUES (?, ?, ?)`)

	defer wordSt.Close()
	defer translationSt.Close()
	defer phraseSt.Close()

	for _, word := range words {
		//插入单词表
		res, err := wordSt.Exec(word.Word)
		if err != nil {
			log.Println("插入单词表失败:", err)
			continue
		}
		var wordID int64
		wordID, err = res.LastInsertId()
		if wordID == 0 {
			// 如果是已存在的 word，查询它的 id
			err = tx.QueryRow(`SELECT id FROM words WHERE word = ?`, word.Word).Scan(&wordID)
			if err != nil {
				log.Println("查询 word_id 失败:", err)
				continue
			}
		}

		//插入词性表
		for _, translation := range word.Translations {
			_, err := translationSt.Exec(wordID, translation.Type, translation.Translation)
			if err != nil {
				log.Println("插入词性表失败:", err)
			}
		}

		//插入短语表
		for _, phrase := range word.Phrases {
			_, err := phraseSt.Exec(wordID, phrase.Phrase, phrase.Translation)
			if err != nil {
				log.Println("插入短语表失败:", err)
			}
		}
	}

	//提交事务
	if err := tx.Commit(); err != nil {
		log.Fatal("提交事务失败:", err)
	}
}

func main() {
	//记录开始时间
	start := time.Now()

	files := []string{
		"json/english-vocabulary/json/3-CET4-顺序.json",
		"json/english-vocabulary/json/4-CET6-顺序.json",
	}

	//打开并读取文件内容
	var allWords []Word
	for _, path := range files {
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("打开文件失败：%s,错误：%v", path, err)
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal("读取文件失败:", err)
		}
		//解析JSON内容
		var words []Word
		if err := json.Unmarshal(data, &words); err != nil {
			log.Fatal("解析JSON失败:", err)
		}
		allWords = append(allWords, words...) // 合并所有单词
	}
	start1 := time.Since(start).Seconds()
	fmt.Printf("读取文件完成,耗时：%.2f\n", start1)

	//连接数据库
	db, err := sql.Open("sqlite", "words.db")
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	defer db.Close()

	//创建表
	createTables(db)
	insertData(db, allWords)

	//记录结束时间并计算耗时
	elapsed := time.Since(start).Seconds()
	fmt.Printf("数据导入完成,耗时：%.2f\n", elapsed)
}
