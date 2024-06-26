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

func LiveHandler(db *gorm.DB, topic string, msg string, devices []models.TimeoutDevices) (string, []models.TimeoutDevices) {
	liveData := viewModels.LiveViewModel{}

	err := json.Unmarshal([]byte(msg), &liveData)

	if err != nil {
		log.Printf("Failed to parse message: %v", err)
		return fmt.Sprintf("ERROR: Failed to parse message"), nil
	}

	var deviceTable models.Device
	if err := db.Where(&models.Device{SerialNumber: liveData.SerialNumber}).First(&deviceTable).Error; err != nil {
		log.Printf("Serial number not found in Devices table: %v", err)
		return fmt.Sprintf("ERROR: Serial number not found in Devices table"), nil
	}

	if !utils.IsValidStatus(liveData.Status) {
		log.Printf("Invalid status: %v", liveData.Status)
		return fmt.Sprintf("ERROR: Invalid status"), nil
	}
	for _, device := range devices {
		if device.SerialNumber == liveData.SerialNumber {
			device.Counter = 0
		}
	}
	return "OK: You are a live! :)", devices
}
