package syntax

import (
	"github.com/PoiCraft/PoiCraftQQ-V3/command"
	"strings"
)

var (
	Prefix string
)

func GroupMsg(msg string, fromQQ int64, ret func(msg string)) bool {
	if strings.HasPrefix(msg, Prefix) {
		cmd := strings.TrimPrefix(msg, Prefix)
		args := strings.Fields(cmd)
		if len(args) < 1 {
			return false
		}
		switch args[0] {
		case "bind":
			command.Bind(args, cmd, fromQQ, ret)
		case "randomtp":
			command.RandomTp(args, cmd, fromQQ, ret)
		}
	}
	return false
}
