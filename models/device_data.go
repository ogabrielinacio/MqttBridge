package models

import (
	"github.com/google/uuid"
	"mqttbridge/enums"
	"time"
)

type DeviceData struct {
	DevicesDataID uuid.UUID `gorm:"primaryKey"`
	SerialNumber  string
	Status        enums.EStatus 
	SoilMoisture  int
	Date          time.Time
}

func (DeviceData) TableName() string {
	return "DeviceData"
}