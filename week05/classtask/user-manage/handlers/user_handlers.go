package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age"`
}

func GetUsers(c *gin.Context) {
	data, err := ioutil.ReadFile("user.json")
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusOK, []User{})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "读取文件失败"})
		return
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解析文件内容失败"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// 创建新用户
func CreateUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式错误"})
		return
	}

	//读取现有用户
	existingUsers, err := getExistingUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "读取现有用户失败"})
		return
	}

	//检查邮箱是否唯一
	for _, user := range existingUsers {
		if user.Email == newUser.Email {
			c.JSON(http.StatusBadRequest, gin.H{"error": "该邮箱已存在"})
			return
		}
	}

	existingUsers = append(existingUsers, newUser)
	err = saveUsersToFile(existingUsers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存用户失败"})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

// 更新用户
func UpdateUser(c *gin.Context) {
	var UpdateUser User
	if err := c.ShouldBindJSON(&UpdateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式错误"})
		return
	}

	existingUsers, err := getExistingUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "读取现有用户失败"})
		return
	}

	found := false
	for i, user := range existingUsers {
		if user.Email == UpdateUser.Email {
			existingUsers[i] = UpdateUser
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}
	err = saveUsersToFile(existingUsers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存更新后的用户失败"})
		return
	}
	c.JSON(http.StatusOK, UpdateUser)
}

func DeleteUser(c *gin.Context) {
	email := c.Param("email")
	existingUsers, err := getExistingUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "读取现有用户失败"})
		return
	}

	newUsers := make([]User, 0, len(existingUsers))
	for _, user := range existingUsers {
		if user.Email != email {
			newUsers = append(newUsers, user)
		}
	}
	if len(newUsers) == len(existingUsers) {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}
	err = saveUsersToFile(newUsers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存删除后的用户列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
}

func getExistingUsers() ([]User, error) {
	data, err := ioutil.ReadFile("user.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []User{}, nil
		}
		return nil, err
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func saveUsersToFile(users []User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile("user.json", data, 0644)
}
