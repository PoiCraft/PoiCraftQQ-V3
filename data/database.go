package data

import (
	"github.com/PoiCraft/PoiCraftQQ-V3/models"
	"github.com/jinzhu/gorm"
)

func InitDB() {
	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open("mysql", Conf.MySQL)
	if err != nil {
		panic("数据库连接不成功：" + err.Error())
	}
	models.DB = db
	models.Migrate()
}
