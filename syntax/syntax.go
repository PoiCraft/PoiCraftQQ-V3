package syntax

import (
	"github.com/PoiCraft/PoiCraftQQ-V3/whitelist"
	"strings"
)

var (
	Prefix string
)

func GroupMsg(msg string, ret func(msg string)) bool {
	if strings.HasPrefix(msg, Prefix) {
		cmd := strings.TrimPrefix(msg, Prefix)
		args := strings.Fields(cmd)
		if len(args) < 1 {
			return false
		}
		switch args[0] {
		case "whitelist":
			whitelist.Whitelist(args, cmd, ret)
		}
	}
	return false
}
