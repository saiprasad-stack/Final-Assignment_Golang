package controllers

import (
	"net/http"
	"strconv"

	"billing-system/backend/services"

	"github.com/gin-gonic/gin"
)

func GetBill(c *gin.Context) {

	id, _ := strconv.Atoi(c.Query("consumer_id"))
	date := c.Query("date")

	total, breakdown, firstReading, lastReading := services.CalculateBill(uint(id), date)

	c.JSON(http.StatusOK, gin.H{
		"total":         total,
		"breakdown":     breakdown,
		"first_reading": firstReading,
		"last_reading":  lastReading,
	})
}
