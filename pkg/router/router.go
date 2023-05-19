package router

import (
	"fmt"
	"gig-echo-template/pkg/env"
	"gig-echo-template/pkg/templates"
	"gig-echo-template/src/routes"
	"path/filepath"
	"text/template"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Run router
func Run(db *gorm.DB) {
	path, _ := filepath.Abs("./pkg/templates/index.html")
	r := echo.New()
	r.Renderer = &templates.TemplateRenderer{
		Templates: template.Must(template.ParseFiles(path)),
	}
	routes.AppRoutes(r, db)

	port := fmt.Sprintf(":%s", env.Get("APP_PORT", "8080"))

	r.Logger.Fatal(r.Start(port))
}
