package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	gorm.Model
	Email            string `gorm:"unique"`
	FirstName        string
	LastName         string
	Country          string
	PhoneNumber      string `gorm:"unique"`
	BusinessName     string
	BusinessType     string
	LightningAddress string
	Password         string
	Balance          float64
	Role             string
}
