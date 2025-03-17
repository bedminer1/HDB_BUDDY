package main

import (
	"github.com/bedminer1/HDB_BUDDY/v3/internal/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func initHandler() *handler {
	db, err := gorm.Open(sqlite.Open("../../hdb.db"), &gorm.Config{})
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

	// averages and computes time based records
	// change in price per sqft over time

	// ====== WATCHLIST =======

	// ===== LEADERBOARDS ======

	// returns
	// 1. assets owned by user
	// 2. graph data, for scaled total price and price per sqft
	return c.JSON(200, echo.Map{
		"user": user,
	})
}
