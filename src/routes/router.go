package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// List of app routes
func AppRoutes(r *echo.Echo, db *gorm.DB) {
	r.GET("/", func(ctx echo.Context) error {
		return ctx.Render(200, "index.html", map[string]any{})
	})
	// Define your routes group here ...
	apiRoutes(r, db)

}
