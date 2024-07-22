package model

import (
	//"mytechblog/utils"
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