package register

import (
	"gf-mqtt-handler/mqtt/event"
	"gf-mqtt-handler/mqtt/global"
)

// CommandEvent 注册指令
var CommandEvent = map[string]global.EventHandler{
	"test": event.Test, // 测试
}
