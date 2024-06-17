package models

import (
	"github.com/google/uuid"
	"mqttbridge/enums"
	"time"
)

type DeviceData struct {
	DevicesDataID uuid.UUID     `gorm:"column:DeviceDataId;primaryKey"`
	SerialNumber  string        `gorm:"column:SerialNumber"`
	Status        enums.EStatus `gorm:"column:Status"`
	SoilMoisture  int           `gorm:"column:SoilMoisture"`
	Date          time.Time     `gorm:"column:Date"`
}

func (DeviceData) TableName() string {
	return "DeviceData"
}
