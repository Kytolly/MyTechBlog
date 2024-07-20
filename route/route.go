package route

import(
	"log/slog"
	"mytechblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(){
	slog.Info("Initializing Router...")

	gin.SetMode(utils.AppMode)

	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"msg": "hello",
			})
		})
	}
	slog.Info("The Project Is initted on http://localhost:4040/api/v1/hello !")
	r.Run(utils.HttpPort)
}