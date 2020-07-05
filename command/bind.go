package command

import (
	"fmt"
	"github.com/PoiCraft/PoiCraftQQ-V3/models"
	"strconv"
	"strings"
)

func Bind(args []string, cmd string, fromQQ int64, ret func(msg string)) {
	if len(args) < 2 {
		ret("您的指令不合法哦 发送#help获取使用帮助")
		return
	}
	gamerName := strings.Join(args[1:], " ")
	qqNumber := strconv.FormatInt(fromQQ, 10)
	if models.IsUserBindXbox(qqNumber) {
		ret(fmt.Sprintf("您已经绑定过Xbox账号了，不能再绑定了，如需解绑请发送 #unbind"))
		return
	}

	if models.IsGamerBindQQ(gamerName) {
		ret(fmt.Sprintf("您的Xbox账号已经被别人绑定了哦，如需解绑请在游戏里发送 #unbind"))
		return
	}

	user := models.User{
		GamerName: gamerName,
		QQNumber:  qqNumber,
		Status:    models.Active,
	}

	models.DB.Create(&user)

	ret("已经成功绑定！")
}
