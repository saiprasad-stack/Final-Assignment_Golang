package controllers

import (
	"billing-system/backend/config"
	"billing-system/backend/models"

	"github.com/gin-gonic/gin"
)

func GetLocations(c *gin.Context) {
	var locations []models.Location
	config.DB.Find(&locations)
	c.JSON(200, locations)
}
