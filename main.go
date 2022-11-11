package main

import (
	"gf-mqtt-handler/mqtt/client"
	"gogs.mirlowz.com/x/gf-x-mqtt/mqtt"
)

func main() {
	mqtt.CreateClient().SetMessageCallbackFunc(client.MessageHandler).Run()
}
