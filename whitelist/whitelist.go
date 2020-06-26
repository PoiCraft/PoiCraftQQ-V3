package whitelist

func Whitelist(args []string, cmd string, ret func(msg string)) {
	if len(args) < 2 {
		ret("您的指令不合法哦 发送#help获取使用帮助")
		return
	}

	switch args[1] {
	case "add":
		add(args, cmd, ret)
	}
}
