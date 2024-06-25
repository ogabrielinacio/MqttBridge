package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"mqttbridge/models"
	"mqttbridge/utils"
	"mqttbridge/viewModels"
	"time"
)

func DataHandler(db *gorm.DB, topic string, msg string) string {
	newData := viewModels.DeviceDataViewModel{}

	err := json.Unmarshal([]byte(msg), &newData)

	if err != nil {
		log.Printf("Failed to parse message: %v", err)
		return fmt.Sprintf("ERROR: Failed to parse message")
	}

	var deviceTable models.Device
	if err := db.Where(&models.Device{SerialNumber: newData.SerialNumber}).First(&deviceTable).Error; err != nil {
		log.Printf("Serial number not found in Devices table: %v", err)
		return fmt.Sprintf("ERROR: Serial number not found in Devices table")
	}

	if !utils.IsValidStatus(newData.Status) {
		log.Printf("Invalid status: %v", newData.Status)
		return fmt.Sprintf("ERROR: Invalid status")
	}

	var deviceData = models.DeviceData{
		DevicesDataID: uuid.New(),
		SerialNumber:  newData.SerialNumber,
		Status:        newData.Status,
		SoilMoisture:  newData.SoilMoisture,
		Date:          time.Now().UTC(),
	}

	err = db.Create(&deviceData).Error
	if err != nil {
		log.Printf("Failed to insert device into the database: %v", err)
		return fmt.Sprintf("ERROR: Failed to insert device into the database")
	}

	log.Printf("New device registered: %+v", newData)
	return fmt.Sprintf("OK: New device registered %+v", newData)
}
