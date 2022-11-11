package client

import (
	"fmt"
	"gf-mqtt-handler/mqtt/event"
	"gf-mqtt-handler/mqtt/global"
	oMqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	xMqtt "gogs.mirlowz.com/x/gf-x-mqtt/mqtt"
)

// CommandEvent 注册指令
var CommandEvent = map[string]global.EventHandler{
	"test": event.Test, // 测试
}

var Client *xMqtt.Client

func MessageHandler(client *xMqtt.Client, client2 oMqtt.Client, message oMqtt.Message) {
	msg := message.Payload()
	// 解析接收数据
	dataJson, err := gjson.DecodeToJson(msg)

	data := dataJson.GetJson("data")
	// 获取 ctx
	ctx := gctx.New()
	// 处理 json 解析异常
	if err != nil {
		g.Log().Error(ctx, fmt.Sprintf("MQTT JSON 解析出错 接收到的数据: %s", msg), err.Error())
		return
	}
	// 处理事件
	Command := dataJson.Get("command").String()

	if fn := CommandEvent[Command]; fn != nil {
		fmt.Println(global.EventData{
			SendTopic: message.Topic(),
			Command:   Command,
			MessageId: message.MessageID(),
			Data:      data,
			Ctx:       ctx,
			Client:    client,
		})

		fn(&global.EventData{
			SendTopic: message.Topic(),
			Command:   Command,
			MessageId: message.MessageID(),
			Data:      data,
			Ctx:       ctx,
			Client:    client,
		})
	}
}
