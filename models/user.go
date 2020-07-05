package models

import "github.com/jinzhu/gorm"

const (
	// 需要QQ 验证 Value0
	NeedQQValidate = iota
	// 需要去游戏里验证 Value1
	NeedXboxValidate
	// 已验证通过 Value2
	Active
)

type User struct {
	gorm.Model
	GamerName string
	QQNumber  string
	Status    int
}

func IsUserBindXbox(QQNumber string) bool {
	var user User
	DB.Where("qq_number = ?", QQNumber).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func IsGamerBindQQ(GamerName string) bool {
	var user User
	DB.Where("gamer_name = ?", GamerName).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
