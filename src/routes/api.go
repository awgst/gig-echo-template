package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Api routes list
func apiRoutes(r *echo.Echo, db *gorm.DB) {
	api := r.Group("/api")
	api.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "pong",
		})
	})
	// Define your routes here
}
