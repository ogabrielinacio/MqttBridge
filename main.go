package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"mqttbridge/handlers"
	"mqttbridge/models"
	"mqttbridge/utils"
	"regexp"
	"time"
)

func main() {
	dbConfig := utils.GetConnectionString()

	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	mqttCfg := utils.GetMqttUtils()
	clientID := uuid.New().String()

	opts := mqtt.NewClientOptions().
		AddBroker(fmt.Sprintf("tcp://%s:%d", mqttCfg.Host, mqttCfg.Port)).
		SetClientID(clientID).
		SetUsername(mqttCfg.Username).
		SetPassword(mqttCfg.Password)

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to the MQTT broker: %v", token.Error())
	}

	regex := regexp.MustCompile(`ERROR|OK`)
	client.Subscribe("register", 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())
		message := string(msg.Payload())
		if !regex.MatchString(message) {
			result := handlers.RegisterHandler(db, msg.Topic(), message)
			client.Publish(msg.Topic(), 0, false, result)
		}
	})
	client.Subscribe("data", 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())
		message := string(msg.Payload())
		if !regex.MatchString(message) {
			result := handlers.DataHandler(db, msg.Topic(), message)
			client.Publish(msg.Topic(), 0, false, result)
		}
	})

	var timeoutDevicesArray []models.TimeoutDevices

	client.Subscribe("live", 0, func(client mqtt.Client, msg mqtt.Message) {
		log.Printf("Received message: %s from topic: %s", msg.Payload(), msg.Topic())
		message := string(msg.Payload())
		if !regex.MatchString(message) {
			resultMessage, _ := handlers.LiveHandler(db, msg.Topic(), message, timeoutDevicesArray)
			client.Publish(msg.Topic(), 0, false, resultMessage)
		}
	})

	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				timeoutDevices := handlers.TimeoutLiveHandler(db)
				for _, device := range timeoutDevices {
					if len(timeoutDevicesArray) == 0 {
						timeoutDevicesArray = append(timeoutDevicesArray, models.TimeoutDevices{SerialNumber: device.SerialNumber, Counter: 0})
					}
					for i, timeoutDevice := range timeoutDevicesArray {
						if timeoutDevice.SerialNumber == device.SerialNumber {
							if timeoutDevicesArray[i].Counter < 3 {
								timeoutDevicesArray[i].Counter++
							}
							if timeoutDevicesArray[i].Counter >= 3 {
								handlers.OfflineHandler(db, device)
							}
						} else {
							timeoutDevicesArray = append(timeoutDevicesArray, models.TimeoutDevices{SerialNumber: device.SerialNumber, Counter: 0})
						}
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	select {}
}
