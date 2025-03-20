package main

import (
	"strings"

	"github.com/bedminer1/HDB_BUDDY/v3/internal/graph"
	"github.com/bedminer1/HDB_BUDDY/v3/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func initHandler() *handler {
	db, err := gorm.Open(sqlite.Open("./hdb_small.db"), &gorm.Config{})
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
	graphData := graph.RecordsGraph(hdbRecords, user)

	// ====== WATCHLIST =======
	// get list of watchlisted assets from users db
	watchlistedAssets := []models.Asset{}
	watchlistStrSplit := strings.Split(user.Watchlisted, ",")
	for _, assetStr := range watchlistStrSplit {
		assetStrSplit := strings.Split(assetStr, "+")
		watchlistedAssets = append(watchlistedAssets, models.Asset{
			Town:       assetStrSplit[0],
			FlatType:   assetStrSplit[1],
			StreetName: assetStrSplit[2],
		})
	}

	// fetch watchlisted assets from records db
	watchlistGraphData := [][]models.GraphData{}
	for _, asset := range watchlistedAssets {
		assetRecords := []models.HDBRecord{}
		h.DB.Where("town = ? AND street_name = ? AND flat_type = ?", asset.Town, asset.StreetName, asset.FlatType).Find(&assetRecords)
		assetGraphData := graph.RecordsGraph(assetRecords, user)
		watchlistGraphData = append(watchlistGraphData, assetGraphData)
	}

	// ===== LEADERBOARDS ======

	// returns
	// 1. assets owned by user
	// 2. graph data, for scaled total price and price per sqft
	return c.JSON(200, echo.Map{
		"user":            user,
		"graphDataPoints": graphData,
		"watchlistGraphDataPoints": watchlistGraphData,
	})
}
