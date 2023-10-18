package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primary_key;unique"`
	Type          string
	Amount        float64
	Status        string
	CustomerEmail string
	CustomerName  string
	PaymentMethod string
	UserID        string
}
