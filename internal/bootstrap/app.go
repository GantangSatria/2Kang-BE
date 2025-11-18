package bootstrap

import (
	"2Kang/config"
	"2Kang/internal/domain/entity"
	"2Kang/internal/routes"

	"github.com/gofiber/fiber/v3"
)

func InitializeApp() *fiber.App {
    app := fiber.New()

	config.NewPostgres()
	db := config.DB

	// Migrate
	db.AutoMigrate(
		&entity.User{},
		&entity.Tukang{},
	)

	// Register routes
	routeSetup := routes.NewRouteSetup()
	routeSetup.Setup(app)

	return app
}
