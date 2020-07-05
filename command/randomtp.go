package command

import (
	"github.com/PoiCraft/PoiCraftQQ-V3/models"
	"strconv"
)

func RandomTp(args []string, cmd string, fromQQ int64, ret func(msg string)) {
	user, err := models.GetUserByQQ(strconv.FormatInt(fromQQ, 10))
	if err != nil {
		Logger.Waringf("RandomTp Err: %s", err.Error())
		ret("程序内部出现错误，请联系开发者")
		return
	}
	if user.IsAnonymous() {
		ret("你还没有绑定Xbox账号哦吗，请绑定后再使用此功能")
		return
	}
}
