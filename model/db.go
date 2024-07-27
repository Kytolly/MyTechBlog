package model

import (
	"fmt"
	"log/slog"
	"mytechblog/utils"
	"time"
	"os"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func InitDB(){
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser, 
		utils.DbPassword, 
		utils.DbHost, 
		utils.DbPort, 
		utils.DbName,
	)
	slog.Debug(dsn)
	// mytechblog:xqy050116@tcp(localhost:3306)/mytechblog?charset=utf8mb4&parseTime=True&loc=Local

	NewsqlDB := mysql.Open(dsn)
	// if err != nil {
	// 	slog.Error("Failed To Open MySQL database", "error",  err)
	// }
	db, err = gorm.Open(NewsqlDB, &gorm.Config{
		// gorm日志模式：silent
		// Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		// DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		// SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: false,
		},
	})
	if err != nil {
		slog.Error("Failed To Initialize A GormDB", "error", err)
		os.Exit(1)
	}
  	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	slog.Error("Failed to open database", "error", err)
	// }

	// 自动迁移
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	// 禁用默认表明的复数形式
	// db.SingularTable(true)

	sqlDB,_ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10*time.Second)

}