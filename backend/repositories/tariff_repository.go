package repositories

import (
	"billing-system/backend/config"
	"billing-system/backend/models"
)

func GetTariffs(locationID uint) []models.Tariff {
	var t []models.Tariff

	config.DB.Where("location_id = ?", locationID).
		Order("min_units ASC").
		Find(&t)

	return t
}
