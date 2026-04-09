package controllers

import (
	"billing-system/backend/config"
	"billing-system/backend/models"

	"github.com/gin-gonic/gin"
)

func GetConsumers(c *gin.Context) {
	locationID := c.Query("location_id")

	var consumers []models.Consumer
	config.DB.Where("location_id = ?", locationID).Find(&consumers)

	c.JSON(200, consumers)
}
