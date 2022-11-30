package event

import (
	"fmt"
	"gf-mqtt-handler/mqtt/global"
	"github.com/gogf/gf/v2/frame/g"
)

func Test(e *global.EventData) {
	for i := 0; i < 10000; i++ {
		e.SendToCall("testOk", g.Map{
			"a": fmt.Sprintf("copy %d", i+1),
		})
	}
}
