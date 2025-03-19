package main

import (
	"time"

	"github.com/bedminer1/HDB_BUDDY/v3/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func initHandler() *handler {
	db, err := gorm.Open(sqlite.Open("../../hdb_small.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.HDBRecord{}, &models.User{})

	return &handler{DB: db}
}

func (h *handler) handleHealthCheck(c echo.Context) error {
	return c.JSON(200, echo.Map{
		"status": "available",
		"system_info": map[string]string{
			"version": version,
		},
	})
}

func sortRecordsByDate(records []models.HDBRecord) {
	for i := 0; i < len(records)-1; i++ {
		for j := i + 1; j < len(records); j++ {
			if records[i].Time.After(records[j].Time) {
				records[i], records[j] = records[j], records[i]
			}
		}
	}
}

func (h *handler) handleFrontPage(c echo.Context) error {

	// ====== USER STATS =======
	// receive user name
	username := c.QueryParam("username")

	// (MOCKED) query users db,
	// user data should contain assets owned and mortgage
	var user models.User
	if err := h.DB.Table("users").Where("username = ?", username).First(&user).Error; err != nil {
		return c.JSON(200, echo.Map{
			"error": "user not found",
		})
	}

	// queries HDB records db,
	// finding records with matching town n flat type
	hdbRecords := []models.HDBRecord{}
	h.DB.Where("town = ? AND street_name = ? AND flat_type = ?", user.Town, user.StreetName, user.FlatType).Find(&hdbRecords)

	// converts to time based records
	// change in price per sqft over time (daily)
	sortRecordsByDate(hdbRecords)
	graphData := []models.GraphData{}
	startDate := hdbRecords[0].Time
	endDate := hdbRecords[len(hdbRecords)-1].Time

	currentDate := startDate
	lastPrice := 0.0
	lastScaledPrice := 0.0
	recordIndex := 0

	for !currentDate.After(endDate) {
		for recordIndex < len(hdbRecords) && hdbRecords[recordIndex].Time.Truncate(24*time.Hour).Equal(currentDate) {
			lastPrice = hdbRecords[recordIndex].PricePerArea
			lastScaledPrice = lastPrice * float64(user.FloorArea)
			recordIndex++
		}

		graphData = append(graphData, models.GraphData{
			PricePerArea:      lastPrice,
			ScaledResalePrice: lastScaledPrice,
			Date:              currentDate,
		})

		currentDate = currentDate.AddDate(0, 0, 1)
	}

	// ====== WATCHLIST =======

	// ===== LEADERBOARDS ======

	// returns
	// 1. assets owned by user
	// 2. graph data, for scaled total price and price per sqft
	return c.JSON(200, echo.Map{
		"user":            user,
		"graphDataPoints": graphData,
	})
}
