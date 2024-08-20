package route

import (
	"log/slog"
	"mytechblog/mid"
	"mytechblog/utils"

	"github.com/gin-gonic/gin"

	//"net/http"
	"mytechblog/api/v1"
)

func InitRouter(){
	slog.Info("Initializing Router...")

	gin.SetMode(utils.AppMode)

	//r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(mid.GinLogger())
	
	admin_v1 := r.Group("api/v1")
	admin_v1.Use(mid.JwtToken())
	{
		// user 模块的路由接口
		admin_v1.PUT("user/:id", v1.UpdateUser)
		admin_v1.DELETE("user/:id", v1.DeleteUser)
		// category 模块的路由接口
		admin_v1.POST("category/add", v1.AddCategory)
		admin_v1.PUT("category/:id", v1.UpdateCategory)
		admin_v1.DELETE("category/:id", v1.DeleteCategory)
		// article 模块的路由接口
		admin_v1.POST("article/add", v1.AddArticle)
		admin_v1.PUT("article/:id", v1.UpdateArticle)
		admin_v1.DELETE("article/:id", v1.DeleteArticle)
		// 上传文件
		admin_v1.POST("upload", v1.Upload)
	}
	router_v1 := r.Group("api/v1")
	{
		router_v1.POST("user/add", v1.AddUser)
		router_v1.GET("users", v1.GetUsers)
		router_v1.GET("categorys", v1.GetCategorys)
		router_v1.GET("articles", v1.GetArticles)
		router_v1.GET("article/info/:id", v1.GetArticleInfo)
		router_v1.GET("articles/list/:id", v1.GetArticle_Category)
		router_v1.POST("login", v1.Login)
	}
	slog.Info("The Project Is initted on http://localhost:4040/api/v1/hello !")
	r.Run(utils.HttpPort)
}