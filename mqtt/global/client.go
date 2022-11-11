package global

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/text/gstr"
	xMqtt "gogs.mirlowz.com/x/gf-x-mqtt/mqtt"
)

type EventData struct {
	SendTopic string
	Command   string
	MessageId uint16
	Data      *gjson.Json
	Ctx       context.Context
	Client    *xMqtt.Client
}

type EventHandler func(e *EventData)

// SendToCall 向发起用户发送
func (t *EventData) SendToCall(command string, data any) {
	t.Send(command, t.SendTopic, data)
}

// Send 执行发送
func (t *EventData) Send(command string, to string, data any) {
	s := gstr.Split(to, "/")

	t.Client.SendMsg(struct {
		Command string `json:"command"`
		Data    any    `json:"data"`
	}{
		command,
		data,
	}, fmt.Sprintf("%s/r-%s", s[0], s[1]))
}
