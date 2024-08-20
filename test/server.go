package test

import (
	//"log"
	//"log/slog"
	"mytechblog/mid"
	"mytechblog/route"
	"mytechblog/model"
	// "github.com/gin-gonic/gin"
)

func Server(){
	// 设置日志
	mid.SetLogger()

	// 引用数据库
	model.InitDB()

	// 设置路由
    route.InitRouter()
}