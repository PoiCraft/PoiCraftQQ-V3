package whitelist

import "strings"

func add(args []string, cmd string, ret func(msg string)) {
	if len(args) < 3 {
		ret("您的指令不合法哦 发送#help获取使用帮助")
		return
	}
	msg := args[2:]

	ret(strings.Join(msg, " "))
}
