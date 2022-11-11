package event

import (
	"fmt"
	"gf-mqtt-handler/mqtt/global"
	"github.com/gogf/gf/v2/frame/g"
)

func Test(e *global.EventData) {
	fmt.Println(e.Data.Map())

	e.SendToCall("testOk", g.Map{
		"a": "copy",
	})
}
