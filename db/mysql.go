package db

import (
	"calendly/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitMysql() {
	_config := config.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		_config.GetString("db.username"),
		_config.GetString("db.password"),
		_config.GetString("db.uri"),
		_config.GetString("db.name"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_db = db
}

func GetMysql() *gorm.DB {
	return _db
}
