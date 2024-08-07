package model

import (
	"encoding/base64"
	"log/slog"
	msg "mytechblog/utils/errormsg"

	//"fmt"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"log/slog"
	"golang.org/x/crypto/scrypt"
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
	data.Password = ScryptPassword(data.Password)
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
	var err error
	if pageSize > 0 && pageNumber > 0{
		err = db.Limit(pageSize).Offset((pageNumber-1)*pageSize).Find(&users).Error
	}else{
		err = db.Limit(-1).Offset(-1).Find(&users).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound{
		slog.Info("Failed to get users", "error", err)
		return nil
	}
	return users
}

// 编辑用户,不包括密码
func EditUser(id int, data *User)int{
	var user User
	var mp = make(map[string]interface{})
	mp["username"] = data.Username
	mp["role"] = data.Role
	err :=db.Model(&user).Where("id=?", id).Updates(mp).Error
	if err != nil{
		slog.Info("Failed to update user", "error", err)
		return msg.ERROR
	}
	return msg.SUCCESS
}
// 删除用户
func DeleteUser(id int) int{
	var user User
	err := db.Where("id=?", id).Delete(&user).Error
	if err != nil{
		slog.Info("Failed to delete user", "error", err)
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 密码加密
func ScryptPassword(password string)string{
	const Keylen = 10
	salt := []byte{2, 3, 5, 7, 11, 13, 17, 19}

	HashPassword, err :=scrypt.Key([]byte(password), salt, 1<<15, 8, 1, Keylen)
	if err != nil {
		slog.Error("Failed to encrypt password", "error", err)
	}
	FinalPassword := base64.StdEncoding.EncodeToString(HashPassword)
	return FinalPassword
}