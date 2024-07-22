package test

import (
	//"log"
	//"log/slog"
	"mytechblog/logger"
	"mytechblog/route"
	"mytechblog/model"
	// "github.com/gin-gonic/gin"
)

func Server(){
	// 设置日志
	logger.SetLogger()

	// 引用数据库
	model.InitDB()

	// 设置路由
    route.InitRouter()
}