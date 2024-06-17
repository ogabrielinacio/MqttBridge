package handlers

import (
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"mqttbridge/models"
	"mqttbridge/viewModels"
	"time"
)

func DataHandler(db *gorm.DB, topic string, msg string) {
	log.Printf("Received message: %s from topic: %s", msg, topic)

	newData := viewModels.DeviceData{}

	err := json.Unmarshal([]byte(msg), &newData)

	if err != nil {
		log.Printf("Failed to parse message: %v", err)
		return
	}

	var deviceData = models.DeviceData{
		SerialNumber: newData.SerialNumber,
		Status:       newData.Status,
		SoilMoisture: newData.SoilMoisture,
		Date:         time.Now().UTC(),
	}

	err = db.Create(&deviceData).Error
	if err != nil {
		log.Printf("Failed to insert device into the database: %v", err)
		return
	}

	log.Printf("New device registered: %+v", newData)
}
