package repositories

import (
	"billing-system/backend/config"
	"billing-system/backend/models"
)

func GetConsumer(id uint) models.Consumer {
	var c models.Consumer
	config.DB.First(&c, id)
	return c
}
func GetMeterByConsumer(consumerID uint) models.Meter {
	var meter models.Meter
	config.DB.Where("consumer_id = ?", consumerID).First(&meter)
	return meter
}
