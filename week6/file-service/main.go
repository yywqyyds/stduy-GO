package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

type FileMeta struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UUID      string    `json:"uuid"`
	FileName  string    `json:"filename"`
	Path      string    `json:"path"`
	Size      int64     `json:"size"`
	MimeType  string    `json:"mime_type"`
	FileType  string    `json:"file_type"`
	CreatedAt time.Time `json:"created_at"`
}

// 文件表
func createFile(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid TEXT NOT NULL UNIQUE,
		filename TEXT NOT NULL,
		path TEXT NOT NULL UNIQUE,
		size INTEGER NOT NULL,
		mime_type TEXT,
		file_type TEXT,
		uploaded_by TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`)
	if err != nil {
		log.Fatal("创建文件表失败:", err)
	}
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite", "filemeta.db")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	createFile(db)
}

// 上传接口
func uploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil || form.File["files"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": -1,
			"msg":  "没收到文件",
			"data": "",
		})
		return
	}
	allowedFiles := map[string]bool{"jpg": true, "png": true, "js": true, "css": true, "html": true}
	saveDir := "upload" //上传文件夹
	os.MkdirAll(saveDir, os.ModePerm)

	results := []FileMeta{}
	for _, file := range form.File["files"] {
		if file.Size > 5<<20 {
			continue // > 5MB
		}
		ext := strings.ToLower(filepath.Ext(file.Filename))[1:]
		if !allowedFiles[ext] {
			continue
		}

		uuid := uuid.New().String()
		savePath := filepath.Join(saveDir, uuid+"_"+file.Filename)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			continue
		}
		_, err := db.Exec(`INSERT INTO files (uuid, filename, path, size, mime_type, file_type) VALUES (?, ?, ?, ?, ?, ?)`,
			uuid, file.Filename, savePath, file.Size, file.Header.Get("Content-Type"), ext)
		if err == nil {
			results = append(results, FileMeta{UUID: uuid, FileName: file.Filename, Size: file.Size})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": results,
	})
}

// 下载接口
func downloadFile(c *gin.Context) {
	id := c.Param("id")
	row := db.QueryRow("SELECT filename, path FROM files WHERE uuid = ?", id)
	var filename, path string
	if err := row.Scan(&filename, &path); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -4,
			"msg":  "找不到文件",
			"data": "",
		})
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.File(path)
}

// 预览接口
func previewFile(c *gin.Context) {
	id := c.Param("id")
	row := db.QueryRow("SELECT mime_type, path FROM files WHERE uuid = ?", id)
	var mimeType, path string
	if err := row.Scan(&mimeType, &path); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": -4,
			"msg":  "找不到文件",
			"data": "",
		})
		return
	}
	c.Header("content-type", mimeType)
	file, err := os.Open(path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -5,
			"msg":  "无法打开文件",
			"data": "",
		})
		return
	}
	defer file.Close()
	io.Copy(c.Writer, file)
}

// 文件列表接口
func listFiles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	fileType := c.Query("type")
	offset := (page - 1) * size
	query := "SELECT id, uuid, filename, path, size, mime_type, file_type, created_at FROM files"
	args := []interface{}{}
	if fileType != "" {
		query += " WHERE file_type = ?"
		args = append(args, fileType)
	}
	query += " ORDER BY created_at DESC LIMIT ? OFFSET ?"
	args = append(args, size, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -6,
			"msg":  "查询文件失败",
			"data": "",
		})
		return
	}
	defer rows.Close()

	var files []FileMeta
	for rows.Next() {
		var f FileMeta
		var createdAtStr string
		if err := rows.Scan(&f.ID, &f.UUID, &f.FileName, &f.Path, &f.Size, &f.MimeType, &f.FileType, &createdAtStr); err == nil {
			f.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", createdAtStr)
			files = append(files, f)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": files,
	})
}

// 统计接口
func statFiles(c *gin.Context) {
	typeStat := map[string]map[string]int64{}

	rows, err := db.Query("SELECT file_type, COUNT(*), SUM(size) FROM files GROUP BY file_type")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -7,
			"msg":  "统计文件类型失败",
		})
		return
	}
	for rows.Next() {
		var t string
		var count, total int64
		rows.Scan(&t, &count, &total)
		typeStat[t] = map[string]int64{
			"count": count,
			"size":  total,
		}
	}
	rows.Close()

	var totalCount, totalSize int64
	err = db.QueryRow("SELECT COUNT(*), SUM(size) FROM files").Scan(&totalCount, &totalSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -8,
			"msg":  "统计总量失败",
			"data": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": gin.H{
			"total_files": totalCount,
			"total_size":  totalSize,
			"by_type":     typeStat,
		},
	})
}

func main() {
	initDB()
	os.MkdirAll("upload", os.ModePerm)

	r := gin.Default()
	r.POST("/upload", uploadFile)
	r.GET("/download/:id", downloadFile)
	r.GET("/perview/:id", previewFile)
	r.GET("/files", listFiles)
	r.GET("/stats", statFiles)

	r.Run(":8080")
}
