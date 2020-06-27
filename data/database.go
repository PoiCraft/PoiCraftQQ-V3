package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func initDB() {
	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open("mysql", Conf.MySQL)
	if err != nil {
		panic("数据库连接不成功：" + err.Error())
	}
	DB = db
}
