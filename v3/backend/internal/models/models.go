package models

import (
	"time"
	"github.com/google/uuid"
)

type HDBRecord struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Time              time.Time `gorm:"index" json:"time"`
	Town              string    `gorm:"index" json:"town"`
	FlatType          string    `gorm:"index" json:"flat_type"`
	Block             int       `json:"block"`
	StreetName        string    `json:"street_name"`
	StoreyRange       string    `json:"storey_range"`
	FloorArea         int       `json:"floor_area"`
	FlatModel         string    `json:"flat_model"`
	LeaseCommenceDate int       `json:"lease_commence_date"`
	ResalePrice       int       `json:"resale_price"`
	PricePerArea      float64   `json:"price_per_area"`
}