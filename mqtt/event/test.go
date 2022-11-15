package event

import (
	"gf-mqtt-handler/mqtt/global"
	"github.com/gogf/gf/v2/frame/g"
)

func Test(e *global.EventData) {
	e.SendToCall("testOk", g.Map{
		"a": "copy",
	})
}
