package routes

import (
	"2Kang/internal/handler"
	"2Kang/internal/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type RouteSetup struct {
	// authHandler       *handler.AuthHandler
}

func NewRouteSetup() *RouteSetup {
	// authHandler *handler.AuthHandler
	return &RouteSetup{
		// authHandler: authHandler,
	}
}

func (rs *RouteSetup) Setup(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

	api := app.Group("/api/v1")

	api.Get("/", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

}