package model

import (
	//"mytechblog/utils"
	//"fmt"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log/slog"
	msg "mytechblog/utils/errormsg"
)

type Article struct{
	gorm.Model
	Title 		string 	`gorm:"type:varchar(100);not null" json:"title"`
	Category 	Category`gorm:"foreignkey:Cid"`
	Cid 		int 	`gorm:"type:int;not null" json:"cid"`
	Description string 	`gorm:"type:varchar(200)" json:"description"`
	Content 	string 	`gorm:"type:longtext" json:"content"`
	Img 		string 	`gorm:"type:varchar(100)" json:"img"`
}

// 添加文章
func CreateArticle(data *Article)int{
	err := db.Create(&data).Error
	if err != nil{
		slog.Error("Failed to create article")
		return msg.ERROR
	}
	return msg.SUCCESS
}

// 查询分类下的所有文章 
func GetArticle_Category(id int, pageSize int, pageNumber int)([]Article, int, int64){
	var articles_cate []Article
	var err error
	var total int64
	var res *gorm.DB

	if pageSize > 0 && pageNumber > 0{
		res = db.Preload("Category").Limit(pageSize).Offset((pageNumber-1)*pageSize).Where("cid=?", id).Find(&articles_cate)
	}else{
		res = db.Preload("Category").Limit(-1).Offset(-1).Where("cid=?", id).Find(&articles_cate)
	}
	err = res.Error
	total = res.RowsAffected
	if err != nil{
		return nil, msg.ERROR_CATEGORY_NOT_EXIST, 0
	}
	return articles_cate, msg.SUCCESS, total
}
// 查询单个文章
func GetArticleInfo(id int)(Article, int){
	var article Article 
	err = db.Where("id = ?", id).Preload("Category").First(&article).Error
	if err != nil {
		slog.Info("Failed to get article")
		return article, msg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, msg.SUCCESS
}
//查询文章列表
func GetArticles(pageSize int, pageNumber int)([]Article, int, int64){
	var articles []Article
	var err error
	var total int64
	var res *gorm.DB

	if pageSize > 0 && pageNumber > 0{
		res = db.Preload("Category").Limit(pageSize).Offset((pageNumber-1)*pageSize).Find(&articles)
	}else{
		res = db.Preload("Category").Limit(-1).Offset(-1).Find(&articles)
	}
	err = res.Error
	total = res.RowsAffected
	if err != nil && err != gorm.ErrRecordNotFound{
		slog.Info("Failed to get articles", "error", err)
		return nil, msg.ERROR, 0
	}
	return articles, msg.SUCCESS, total
}

// 编辑文章
func EditArticle(id int, data *Article)int{
	var article Article
	var mp = make(map[string]interface{})
	mp["title"] = data.Title
	mp["cid"] = data.Cid
	mp["description"] = data.Description
	mp["content"] = data.Content
	mp["img"] = data.Img
	err :=db.Model(&article).Where("id=?", id).Updates(mp).Error
	if err != nil{
		slog.Info("Failed to update Article", "error", err)
		return msg.ERROR
	}
	return msg.SUCCESS
}
// 删除文章
func DeleteArticle(id int) int{
	var article Article
	err := db.Where("id=?", id).Delete(&article).Error
	if err != nil{
		slog.Info("Failed to delete Article", "error", err)
		return msg.ERROR
	}
	return msg.SUCCESS
}