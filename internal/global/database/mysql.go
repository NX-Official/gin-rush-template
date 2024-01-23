package database

import (
	"fmt"
	"gin-rush-template/config"
	"gin-rush-template/internal/global/query"
	"gin-rush-template/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Query *query.Query

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get().Mysql.Username,
		config.Get().Mysql.Password,
		config.Get().Mysql.Host,
		config.Get().Mysql.Port,
		config.Get().Mysql.DBName,
	)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	// 自动建表
	err = DB.AutoMigrate(
		model.User{},
	)

	if err != nil {
		panic(err)
	}
	Query = query.Use(DB)
}
