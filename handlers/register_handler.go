package handlers

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"mqttbridge/models"
	"mqttbridge/utils"
	"mqttbridge/viewModels"
)

func RegisterHandler(db *gorm.DB, topic string, msg string) {

	log.Printf("Received message: %s from topic: %s", msg, topic)

	newDevice := viewModels.DeviceRegister{}

	err := json.Unmarshal([]byte(msg), &newDevice)

	if err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	if utils.IsSerialNumberValid(newDevice.SerialNumber) == false {
		log.Println("Serial number is invalid")
		return
	}

	hash, salt, err := utils.CreatePasswordHash(newDevice.Password)
	if err != nil {
		log.Println("Error creating Hash")
	}

	var device = models.Device{
		SerialNumber: newDevice.SerialNumber,
		Password:     hash,
		Salt:         salt,
	}

	err = db.Create(&device).Error
	if err != nil {
		log.Printf("Failed to insert device into the database: %v", err)
		return
	}

	log.Printf("New device registered: %+v", newDevice)
}
