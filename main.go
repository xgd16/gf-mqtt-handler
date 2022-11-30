package main

import (
	"gf-mqtt-handler/mqtt/client"
	"gogs.mirlowz.com/x/gf-x-mqtt/mqtt"
)

func main() {
	mqtt.CreateClient(func(option *mqtt.ClientOption, config *mqtt.Config) {
		option.MessageCallbackFunc = client.MessageHandler
	})
}
