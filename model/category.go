package model

import (
	//"mytechblog/utils"
	//"fmt"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	msg "mytechblog/utils/errormsg"
)

type Category struct{
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	ID   uint 	`gorm:"primary_key;auto_increment" json:"id"`
}

// 查询分类是否已存在
func CheckCategory(name string)int{
	var cates Category
	db.Select("id").Where("name=?", name).First(&cates)
	if cates.ID > 0{
		slog.Warn("Category already exists") 
		return msg.ERROR_CATEGORY_USED
	}
	return msg.SUCCESS
}

// 添加分类
func CreateCategory(data *Category)int{
	err := db.Create(&data).Error
	if err != nil{
		slog.Error("Failed to create Category")
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 查询分类列表
func GetCategorys(pageSize int, pageNumber int)[]Category{
	var cates []Category
	var err error
	if pageSize >0 && pageNumber >0{
		err = db.Limit(pageSize).Offset((pageNumber-1)*pageSize).Find(&cates).Error
	}else{
		err = db.Limit(-1).Offset(-1).Find(&cates).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound{
		slog.Info("Failed to get users", "error", err)
		return nil
	}
	return cates
}

// 编辑分类
func EditCategory(id int, data *Category)int{
	var cate Category
	var mp = make(map[string]interface{})
	mp["name"] = data.Name
	err :=db.Model(&cate).Where("id=?", id).Updates(mp).Error
	if err != nil{
		slog.Info("Failed to update user", "error", err)
		return msg.ERROR
	}
	return msg.SUCCESS
}
// 删除分类
func DeleteCategory(id int) int{
	var cate Category
	err := db.Where("id=?", id).Delete(&cate).Error
	if err != nil{
		slog.Info("Failed to delete user", "error", err)
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 查询分类下的所有文章
