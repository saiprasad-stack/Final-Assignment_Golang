package routes

import (
	"billing-system/backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/locations", controllers.GetLocations)
	r.GET("/consumers", controllers.GetConsumers)
	r.GET("/bill", controllers.GetBill)
}
