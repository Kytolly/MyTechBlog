package mid

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

func CORS() gin.HandlerFunc{
	return func(c *gin.Context){
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			// AllowCredentials: true,
			// AllowOriginFunc: func(origin string) bool {
			//   return origin == "https://github.com"
			// },
			MaxAge: 12 * time.Hour,
		  })
	}
}