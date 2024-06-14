package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `gorm:"primaryKey"`
	Name      string
	LastName  string
	Email     string
	Password  []byte
	Salt      []byte
	Devices   []*Device `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "Users"
}