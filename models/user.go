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
	TpNumber  int
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

func GetUserByQQ(QQNumber string) (User, error) {
	var user User
	result := DB.Where("qq_number = ? AND status = ?", QQNumber, Active).First(&user)
	return user, result.Error
}

func (user *User) IsAnonymous() bool {
	return user.ID == 0
}
