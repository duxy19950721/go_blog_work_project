package config

import (
	"blog/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:12345678@tcp(localhost:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		panic(error)
	}

	// 表结构同步映射
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

	return db
}
