package model

import (
	"log/slog"
	msg "mytechblog/utils/errormsg"
	//"fmt"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"log/slog"
)

type User struct{
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(500);not null" json:"password"`
	Role 	 int 	`gorm:"type:int" json:"role"`
}

// 查询用户是否已存在
func CheckUser(name string)int{
	var users User
	db.Select("id").Where("username=?", name).First(&users)
	if users.ID > 0{
		slog.Warn("User already exists") 
		return msg.ERROR_USERNAME_USED
	}
	return msg.SUCCESS
}

// 添加用户
func CreateUser(data *User)int{
	err := db.Create(&data).Error
	if err != nil{
		slog.Error("Failed to create user")
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 查询用户列表
func GetUsers(pageSize int, pageNumber int)[]User{
	var users []User
	err := db.Limit(pageSize).Offset((pageNumber-1)*pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		slog.Info("Failed to get users", "error", err)
		return nil
	}
	return users
}

// 编辑用户