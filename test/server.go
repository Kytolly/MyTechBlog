package test

import (
	//"log"
	//"log/slog"
	"mytechblog/logger"
	"mytechblog/route"
	// "github.com/gin-gonic/gin"
)

func Server(){
	logger.SetLogger()
    route.InitRouter()
}