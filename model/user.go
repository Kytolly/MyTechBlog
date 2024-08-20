package model

import (
	"encoding/base64"
	"log/slog"
	msg "mytechblog/utils/errormsg"
	"fmt"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"log/slog"
	"golang.org/x/crypto/scrypt"
)

type User struct{
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role 	 int 	`gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"身份码"`
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
func GetUsers(pageSize int, pageNumber int)([]User, int64){
	var users []User
	var err error
	var total int64
	var res *gorm.DB
	if pageSize > 0 && pageNumber > 0{
		res = db.Limit(pageSize).Offset((pageNumber-1)*pageSize).Find(&users)
	}else{
		res = db.Limit(-1).Offset(-1).Find(&users)
	}
	err = res.Error
	total = res.RowsAffected
	slog.Debug(fmt.Sprintf("%d", total))
	if err != nil && err != gorm.ErrRecordNotFound{
		slog.Info("Failed to get users", "error", err)
		return nil, 0
	}
	return users, total
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

// 登录验证
func CheckLogin(username string, password string) (int){
	//slog.Debug(username)
	//slog.Debug(password)
	var user User
	db.Where("username=?", username).First(&user)
	// err := db.Where("username=?", username).First(&user).Error
	// if err != nil{
	// 	return msg.ERROR
	// }
	if user.ID == 0{
		return msg.ERROR_USER_NOT_EXIST
	}
	if ScryptPassword(password) != user.Password{
		return msg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1{
		return msg.ERROR_USER_NO_RIGHT
	}
	return msg.SUCCESS
}