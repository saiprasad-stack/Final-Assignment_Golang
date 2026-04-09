package main

import (
	"billing-system/backend/config"
	"billing-system/backend/models"
	"billing-system/backend/routes"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(
		&models.Consumer{},
		&models.Location{},
		&models.Meter{},
		&models.BillingData{},
		&models.Tariff{},
	)

	r := gin.Default()
	r.Use(cors.Default())
	routes.SetupRoutes(r)

	r.Run(":8080")
}
