package handlers

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"mqttbridge/models"
	"mqttbridge/utils"
	"mqttbridge/viewModels"
)

func RegisterHandler(db *gorm.DB, topic string, msg string) string {

	newDevice := viewModels.DeviceRegister{}

	err := json.Unmarshal([]byte(msg), &newDevice)

	if err != nil {
		log.Println("Failed to parse message: %v", err)
		return fmt.Sprintf("ERROR: Failed to parse message")
	}

	if utils.IsSerialNumberValid(newDevice.SerialNumber) == false {
		log.Println("Serial number is not valid")
		return fmt.Sprintf("ERROR: Serial number is invalid")
	}

	hash, salt, err := utils.CreatePasswordHash(newDevice.Password)
	if err != nil {
		log.Println("Failed creating Hash, %v", err)
		return fmt.Sprintf("ERROR: Failed to create Hash")
	}

	var device = models.Device{
		SerialNumber: newDevice.SerialNumber,
		Password:     hash,
		Salt:         salt,
	}

	err = db.Create(&device).Error
	if err != nil {
		log.Printf("Failed to insert device into the database: %v", err)
		return fmt.Sprintf("ERROR: Failed to insert device into the database")
	}

	log.Printf("New device registered: %+v", newDevice)
	return fmt.Sprintf("OK: New device registered: %+v", newDevice)
}
