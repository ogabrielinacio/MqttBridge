package handlers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"mqttbridge/enums"
	"mqttbridge/models"
	"time"
)

func OfflineHandler(db *gorm.DB, device models.Device) {

	var newDeviceData = models.DeviceData{
		DevicesDataID: uuid.New(),
		SerialNumber:  device.SerialNumber,
		Status:        enums.Offline,
		SoilMoisture:  0,
		Date:          time.Now().UTC(),
	}

	var deviceData models.DeviceData
	err := db.Where(&models.DeviceData{SerialNumber: device.SerialNumber}).Order(models.DeviceData{}.Date).First(&deviceData).Error

	if err != nil {
		log.Printf("Error on Query to get device data for device %s: %v", device.SerialNumber, err)
	}

	if deviceData.Status != enums.Offline {
		err := db.Create(&newDeviceData).Error

		if err != nil {
			log.Printf("Failed to insert device into the database: %v", err)
		}

	}
}
