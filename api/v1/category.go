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

// 查询分类是否已存在
func CategoryExist(c *gin.Context)int{
	return 1
}

// 添加分类
// 不允许相同名字的用户写入
func AddCategory(c *gin.Context){
	slog.Debug(c.FullPath())

	var data model.Category
	err := c.ShouldBindJSON(&data)
	if err != nil{
		slog.Error("Bind JSON Error")
	}
	slog.Debug(data.Name)

	status := model.CheckCategory(data.Name)

	if status == msg.SUCCESS{
		status = model.CreateCategory(&data)
	}
	if status == msg.ERROR_CATEGORY_USED{
		slog.Error("User already exists")
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data": data,
		"message": msg.GetErrorMessage(status),
	})
}

// 查询分类列表
func GetCategorys(c *gin.Context){
	pageSize,_ := strconv.Atoi(c.Query("pagesize"))
	pageNumber,_ := strconv.Atoi(c.Query("pagenumber"))
	data := model.GetCategorys(pageSize, pageNumber)
	status := msg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data": data,
		"message": msg.GetErrorMessage(status),
	})
}
// 更新分类
func UpdateCategory(c *gin.Context){
	var data model.Category
	c.ShouldBindJSON(&data)
	id,_ := strconv.Atoi(c.Param("id"))
	slog.Debug(fmt.Sprintf("%d", id))
	
	status := model.CheckCategory(data.Name)
	if status == msg.SUCCESS{
		status = model.EditCategory(id, &data)
	}
	if status == msg.ERROR_CATEGORY_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"message": msg.GetErrorMessage(status),
	})
}

// 删除分类
func DeleteCategory(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))

	status := model.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"message": msg.GetErrorMessage(status),
	})
}