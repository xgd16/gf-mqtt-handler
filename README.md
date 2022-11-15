# [GF](https://goframe.org/pages/viewpage.action?pageId=1114119) 的 MQTT 事件驱动

## 注意
    客户端发送消息 topic 需要遵循 xx/xx/xx 格式

### 1. 配置
    [mqtt]
    debug = false
    url = "tcp://xxx.xxx.xxx.xxx:1883"
    clientId = "xxx_mqtt_service_client"
    subscribe = "xxx/#"
    qos = 2

### 2. 事件注册
    mqtt/register/event.go
    CommandEvent map 中进行注册
    key 填写事件名 value 事件函数

