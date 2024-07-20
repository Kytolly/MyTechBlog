package logger

import (
	"log/slog"
	"os"
	"fmt"
)

var mapLevelStr = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

var loggerLevel = "debug"
// var loggerfile  = "filelog.log"
var l *slog.Logger

func SetLogger() error{
	Level, ok:= mapLevelStr[loggerLevel]
	if !ok {
		return fmt.Errorf("unknown logger level: %s", loggerLevel)
	}

	// File, err := os.Open(loggerfile)
	// if err != nil {
	// 	return err
	// }
	// defer File.Close()

	Options := &slog.HandlerOptions{
		AddSource:  	true, // 记录日志位置
		Level: 			Level, // 设置日志等级
		ReplaceAttr: 	nil,
	}

	consoleHandler := slog.NewJSONHandler(os.Stdout, Options)
	// fileHandler := slog.NewTextHandler(File, Options)
	l = slog.New(consoleHandler)
	slog.SetDefault(l)

	return nil
}