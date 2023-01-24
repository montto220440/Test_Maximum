package db

import (
	"log"
	"api-pos/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func Connectdb() {
	dsn := "root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err!=nil{
		log.Fatal("connect db failed")
	}
	Conn = db
}

func Migrate(){
	Conn.AutoMigrate(
		&models.Todo_list{},
	)
}