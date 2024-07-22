package route

import(
	"log/slog"
	"mytechblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"mytechblog/api/v1"
)

func InitRouter(){
	slog.Info("Initializing Router...")

	gin.SetMode(utils.AppMode)

	r := gin.Default()
	router_v1 := r.Group("api/v1")
	{
		router_v1.GET("hello", func(c *gin.Context){
			c.JSON(http.StatusOK, gin.H{
				"msg": "hello",
			})
		})

		// user 模块的路由接口
		router_v1.POST("user/add", v1.AddUser)
		router_v1.GET("users", v1.GetUsers)
		router_v1.PUT("user/:id", v1.UpdateUser)
		router_v1.DELETE("user/:id", v1.DeleteUser)
		// category 模块的路由接口


		// 
	}
	slog.Info("The Project Is initted on http://localhost:4040/api/v1/hello !")
	r.Run(utils.HttpPort)
}