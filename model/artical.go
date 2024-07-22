package model

import (
	//"mytechblog/utils"
	//"fmt"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"log/slog"
)

type Artical struct{
	gorm.Model
	Title 		string 	`gorm:"type:varchar(100);not null" json:"title"`
	Category 	Category`gorm:"foreignkey:Cid"`
	Cid 		int 	`gorm:"type:int;not null" json:"cid"`
	Description string 	`gorm:"type:varchar(200)" json:"description"`
	Content 	string 	`gorm:"type:longtext" json:"content"`
	Img 		string 	`gorm:"type:varchar(100)" json:"img"`
}