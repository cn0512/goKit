package ykIoT

/*
	Use IoT.MQTT
*/
import (
	"fmt"
)

type IIoT interface {
	Create(url string)
	PushMsg(buf []byte)
	PopMsg()
}

type YKMqtt struct {
}

func (yk *YKMqtt) Create() {
	fmt.Println("mqtt create")
}
