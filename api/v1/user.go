package v1

import (
	"github.com/gin-gonic/gin"
	msg "mytechblog/utils/errormsg" 
)
// 查询用户是否已存在
func UserExist(c *gin.Context)int{
	return msg.SUCCESS
}
// 添加用户
func AddUser(c *gin.Context){
}
// 查询用户
// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context){
}
// 更新用户
func UpdateUser(c *gin.Context){
}
// 删除用户
func DeleteUser(c *gin.Context){
}