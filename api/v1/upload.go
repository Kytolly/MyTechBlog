package v1

import(
	"github.com/gin-gonic/gin"
	"mytechblog/server"
	msg "mytechblog/utils/errormsg"
	"net/http"
	// "log/slog"
)

func Upload(c *gin.Context){
	file, header, _ := c.Request.FormFile("file")
	// slog.Debug(header.Filename)
	url, status := server.UploadFile(file, header)
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"message": msg.GetErrorMessage(status),
		"url": url,
	})
}