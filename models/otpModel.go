package models

import (
	"time"

	"gorm.io/gorm"
)

type Otp struct {
	gorm.Model
	ID        string `gorm:"unique"`
	Email     string
	Otp       string
	ExpiredAt time.Time
}
