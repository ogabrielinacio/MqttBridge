package utils 

import (
	"os"
	"strconv"
)

type MqttUtils struct {
	Host     string
	Port     int
	Username string
	Password string
}

func GetMqttUtils() MqttUtils {
	var mqttUtils MqttUtils

	mqttUtils.Host = os.Getenv("MQTT_HOST")
	port, err := strconv.Atoi(os.Getenv("MQTT_PORT"))
	if err != nil {
		port = 1883 
	}
	mqttUtils.Port = port
	mqttUtils.Username = os.Getenv("MQTT_USERNAME")
	mqttUtils.Password = os.Getenv("MQTT_PASSWORD")

	return mqttUtils
}