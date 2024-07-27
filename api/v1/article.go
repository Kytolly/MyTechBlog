package v1

import (
	"fmt"
	"log/slog"
	"mytechblog/model"
	msg "mytechblog/utils/errormsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 添加文章
// 不允许相同名字的用户写入
func AddArticle(c *gin.Context){
	slog.Debug(c.FullPath())
	var data model.Article
	err := c.ShouldBindJSON(&data)
	if err != nil{
		slog.Error("Bind JSON Error")
	}
	status := model.CreateArticle(&data)
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data": data,
		"message": msg.GetErrorMessage(status),
	})
}
// 查询单个文章信息
func GetArticleInfo(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	data, status := model.GetArticleInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data": data,
		"message": msg.GetErrorMessage(status),
	})
}
// 查询分类下的所有文章
func GetArticle_Category(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNumber,_ := strconv.Atoi(c.Query("pagenumber"))
	id,_ := strconv.Atoi(c.Param("id"))
	data, status := model.GetArticle_Category(id, pageSize, pageNumber)
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data": data,
		"message": msg.GetErrorMessage(status),
	})
}
// 查询文章列表
func GetArticles(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNumber,_ := strconv.Atoi(c.Query("pagenumber"))
	data, status := model.GetArticles(pageSize, pageNumber)
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data": data,
		"message": msg.GetErrorMessage(status),
	})
}

// 更新文章
func UpdateArticle(c *gin.Context){
	var data model.Article
	c.ShouldBindJSON(&data)
	id,_ := strconv.Atoi(c.Param("id"))
	slog.Debug(fmt.Sprintf("%d", id))
	status := model.EditArticle(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"message": msg.GetErrorMessage(status),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))

	status := model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"message": msg.GetErrorMessage(status),
	})
}