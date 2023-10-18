package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;unique"`
	Title       string
	Description string
	Media       string
	MediaType   string
	Budget      float64
	Status      string
	Views       int64
	Clicks      int64
	Impressions int64
	StartDate   string
	EndDate     string
	Objective   string
	Audience    JSONMap `gorm:"type:jsonb"`
	UserId      string
}
