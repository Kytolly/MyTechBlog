package v1

import (
	"mytechblog/model"
	"mytechblog/mid"
	msg "mytechblog/utils/errormsg"
	"net/http"
	"github.com/gin-gonic/gin"
	//"log/slog"
)

func Login(c *gin.Context){
	var data model.User
	var token string
	var status int
	c.ShouldBindJSON(&data)

	status = model.CheckLogin(data.Username, data.Password)
	if status == msg.SUCCESS{
		token,status  = mid.SetToken(data.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"message": msg.GetErrorMessage(status),
		"token": token,
	})
}