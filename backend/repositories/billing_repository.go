package repositories

import (
	"billing-system/backend/config"
	"billing-system/backend/models"

	"time"
)

func GetMonthlyConsumption(consumerID uint, date string) (float64, float64, float64) {

	selectedDate, _ := time.Parse("2006-01-02", date)
	year, month, _ := selectedDate.Date()

	startOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	selectedDateStr := selectedDate.Format("2006-01-02")
	startOfMonthStr := startOfMonth.Format("2006-01-02")

	var first models.Meter
	config.DB.
		Where("consumer_id = ? AND DATE(date) >= ?", consumerID, startOfMonthStr).
		Order("date ASC").
		First(&first)

	var last models.Meter
	config.DB.
		Where("consumer_id = ? AND DATE(date) <= ?", consumerID, selectedDateStr).
		Order("date DESC").
		First(&last)

	units := last.Reading - first.Reading

	return units, first.Reading, last.Reading
}
