package models

import (
	"github.com/google/uuid"
)

type Device struct {
	SerialNumber string     `gorm:"column:SerialNumber;gorm:primaryKey"`
	Password     []byte     `gorm:"column:Password"`
	Salt         []byte     `gorm:"column:Salt"`
	UserID       *uuid.UUID `gorm:"column:UserId;default:00000000-0000-0000-0000-000000000000"`
}

func (Device) TableName() string {
	return "Devices"
}
