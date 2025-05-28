package orders

import (
	"github.com/arjunksofficial/lumelassignment/internal/entities/orders/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(apiRoutes *gin.RouterGroup) {
	// Create a new h instance
	h := handlers.New()
	orderRoutes := apiRoutes.Group("/orders")
	{
		orderRoutes.GET("/revenue-by-product", h.RevenueByProduct)
		orderRoutes.GET("/revenue-by-category", h.RevenueByCategory)
		orderRoutes.GET("/revenue-by-region", h.RevenueByRegion)
		orderRoutes.GET("/revenue-total", h.RevenueTotal)
		orderRoutes.GET("/revenue-by-trends", h.RevenueByTrends)
	}
}
