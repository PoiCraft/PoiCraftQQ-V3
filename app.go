package main

import (
	"github.com/PoiCraft/PoiCraftQQ-V3/data"
	"github.com/PoiCraft/PoiCraftQQ-V3/syntax"
	"github.com/Tnze/CoolQ-Golang-SDK/cqp"
)

const (
	Ignore    int32 = 0
	Intercept       = 1
)

//go:generate cqcfg -c .
// cqp: 名称: PoiCraft
// cqp: 版本: 3.0.0:1
// cqp: 作者: PoiCraft
// cqp: 简介: PoiCraft 3th QQBot
func main() { /*此处应当留空*/ }

func init() {
	cqp.AppID = "com.poicraft.bot.qq"
	cqp.PrivateMsg = onPrivateMsg
	cqp.GroupMsg = onGroupMsg
	cqp.Start = onStart
}

func onStart() int32 {
	data.Init(cqp.GetAppDir())
	syntax.Prefix = "#"
	return 0
}

func onPrivateMsg(subType, msgID int32, fromQQ int64, msg string, font int32) int32 {
	cqp.SendPrivateMsg(fromQQ, msg) //复读机
	return Ignore
}

func onGroupMsg(subType, msgID int32, fromGroup, fromQQ int64, fromAnonymous, msg string, font int32) int32 {
	ret := func(msg string) {
		cqp.SendGroupMsg(fromGroup, msg)
	}
	syntax.GroupMsg(msg, ret)
	return Ignore
}
