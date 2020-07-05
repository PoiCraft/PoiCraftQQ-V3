package command

import (
	"fmt"
	"github.com/PoiCraft/PoiCraftQQ-V3/data"
	"github.com/PoiCraft/PoiCraftQQ-V3/models"
	"github.com/PoiCraft/PoiCraftQQ-V3/pkg/math"
	"strconv"
	"time"
)

func RandomTp(args []string, cmd string, fromQQ int64, ret func(msg string)) {
	user, err := models.GetUserByQQ(strconv.FormatInt(fromQQ, 10))
	if err != nil {
		ret("你还没有绑定Xbox账号哦吗，请绑定后再使用此功能")
		return
	}

	if user.TpNumber == 0 {
		ret("今日随机传送次数已用完，且行且珍惜，明天再来吧")
		return
	}
	x := strconv.FormatInt(math.RandInt64(30000, 70000, time.Now().UnixNano()-114514+1919-810)*-1, 10)
	z := strconv.FormatInt(math.RandInt64(30000, 70000, time.Now().UnixNano()+810-1919+114514)*-1, 10)

	res, err := data.SendCommand(fmt.Sprintf(`testfor "%s"`, user.GamerName))
	if err != nil {
		Logger.Waringf("RandomTp Execute Command Err: %s", err.Error())
		ret("服务器去火星了，等会儿再试试吧")
		return
	}

	ret(res)

	if res == "No targets matched selector" {
		ret("当前您不在线嗷，上线后再试试吧")
		return
	}

	res, err = data.SendCommand(fmt.Sprintf(`effect "%s" resistance 15 5 true`, user.GamerName))
	if err != nil {
		Logger.Waringf("RandomTp Execute Command Err: %s", err.Error())
		ret("服务器去火星了，等会儿再试试吧")
		return
	}
	ret(res)
	time.Sleep(300 * time.Millisecond)

	res, err = data.SendCommand(fmt.Sprintf(`tp "%s" %s 120 %s`, user.GamerName, x, z))
	ret(fmt.Sprintf(`tp "%s" %s 120 %s`, user.GamerName, x, z))
	if err != nil {
		Logger.Waringf("RandomTp Execute Command Err: %s", err.Error())
		ret("服务器去火星了，等会儿再试试吧")
		return
	}
	ret(res)
	user.TpNumber -= 1
	models.DB.Save(&user)
	ret(fmt.Sprintf(`您已被传送至%s,100,%s
使用次数:%s/%s
传送本不易，且行且珍惜`, x, z, strconv.Itoa(user.TpNumber), strconv.Itoa(data.Conf.TpNumber)))
}
