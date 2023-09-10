package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;unique"`
	gorm.Model
	Email            string `gorm:"unique"`
	FirstName        string
	LastName         string
	Country          string
	PhoneNumber      string
	BusinessName     string
	BusinessType     string
	LightningAddress string
	Password         string
	Balance          float64
	Role             string
}
