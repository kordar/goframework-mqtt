package goframework_mqtt_test

import (
	goframework_mqtt "github.com/kordar/goframework-mqtt"
	logger "github.com/kordar/gologger"
	"testing"
)

func TestMqtt(t *testing.T) {
	cfg := map[string]string{
		"broker":   "192.168.30.104",
		"port":     "1883",
		"username": "",
		"password": "",
	}
	err := goframework_mqtt.AddMqttInstance("cc", cfg)
	logger.Infof("================%v", err)
}
