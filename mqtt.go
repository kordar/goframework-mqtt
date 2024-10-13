package goframework_mqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	logger "github.com/kordar/gologger"
)

type MqttIns struct {
	name string
	ins  mqtt.Client
}

func NewMqttIns(name string, opts *mqtt.ClientOptions) *MqttIns {
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logger.Errorf("[mqtt], err = %v", token.Error())
		return nil
	}
	return &MqttIns{name: name, ins: client}
}

func (c MqttIns) GetName() string {
	return c.name
}

func (c MqttIns) GetInstance() interface{} {
	return c.ins
}

func (c MqttIns) Close() error {
	return nil
}
