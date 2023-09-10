package models

import (
	"time"

	"gorm.io/gorm"
)

type Process struct {
	gorm.Model
	ID      string `gorm:"unique"`
	Email   string
	Process string
	Expiry  time.Time
}

// you can add an expiry to your Process model to check if the processId has expired
