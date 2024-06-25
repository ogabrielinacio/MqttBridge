package handlers

import (
	"gorm.io/gorm"
	"log"
	"mqttbridge/models"
	"time"
)

func TimeoutLiveHandler(db *gorm.DB) []models.Device {
	var devices []models.Device
	result := db.Find(&devices)

	if result.Error != nil {
		log.Printf("Error on Query to get all devices")
	}

	var timeoutDevices []models.Device

	for _, device := range devices {
		var deviceData models.DeviceData
		err := db.Where(&models.DeviceData{SerialNumber: device.SerialNumber}).Order(models.DeviceData{}.Date).Last(&deviceData).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			log.Printf("Error on Query to get device data for device %s: %v", device.SerialNumber, err)
			continue
		}

		if deviceData.Date.Truncate(time.Minute).Add(time.Minute).Before(time.Now().Truncate(time.Minute)) || deviceData.Date.Truncate(time.Minute).Add(time.Minute).Equal(time.Now().Truncate(time.Minute)) {
			timeoutDevices = append(timeoutDevices, device)
		}
	}
	return timeoutDevices
}
