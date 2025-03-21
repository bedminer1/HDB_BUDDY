package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

type User struct {
	gorm.Model
	Username    string   `json:"username"`
	FlatType    string   `gorm:"index" json:"flat_type"`
	Town        string   `gorm:"index" json:"town"`
	StoreyRange string   `json:"storey_range"`
	StreetName  string   `json:"street_name"`
	FloorArea   int      `json:"floor_area"`
	Mortgage    int      `json:"mortgage"`
	Watchlisted string `json:"watchlisted"` // each watchlist item formated town+flatType+streetName
}

type Asset struct {
	Town       string
	FlatType   string
	StreetName string
}

type GraphData struct {
	PricePerArea      float64   `json:"pricePerArea"`
	ScaledResalePrice float64   `json:"resalePrice"`
	Date              time.Time `json:"date"`
}
