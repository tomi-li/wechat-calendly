package db

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitMysql() {
	dsn := "tomi_mysql:Tomi_mysql@tcp(sh-cynosdbmysql-grp-a81aj0gg.sql.tencentcdb.com:26776)/calendly?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	_db = db
}

func GetMysql() *gorm.DB {
	return _db
}
