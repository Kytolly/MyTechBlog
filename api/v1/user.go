package v1

import (
	//"fmt"
	"log/slog"
	"mytechblog/model"
	msg "mytechblog/utils/errormsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 查询用户是否已存在
func UserExist(c *gin.Context)int{
	return 1
}
// 添加用户
// 不允许相同名字的用户写入
func AddUser(c *gin.Context){
	slog.Debug(c.FullPath())

	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil{
		slog.Error("Bind JSON Error")
	}
	slog.Debug(data.Username)

	status := model.CheckUser(data.Username)

	if status == msg.SUCCESS{
		model.CreateUser(&data)
	}
	if status == msg.ERROR_USERNAME_USED{
		slog.Error("User already exists")
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data": data,
		"message": msg.GetErrorMessage(status),
	})
}
// 查询用户
// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNumber,_ := strconv.Atoi(c.Query("pagenumber"))
	if pageNumber == 0{
		pageNumber = 1
	}

	data := model.GetUsers(pageSize, pageNumber)
	status := msg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data": data,
		"message": msg.GetErrorMessage(status),
	})
}
// 更新用户
func UpdateUser(c *gin.Context){
}
// 删除用户
func DeleteUser(c *gin.Context){
}