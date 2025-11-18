package bootstrap

import (
	"2Kang/config"
	"2Kang/internal/domain/entity"
	"2Kang/internal/domain/repository"
	"2Kang/internal/handler"
	"2Kang/internal/routes"
	"2Kang/internal/services"

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

	userRepo := repository.NewUserRepository(config.DB)
	
	authService := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo)
	
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(*userService)

	// Register routes
	routeSetup := routes.NewRouteSetup(
		authHandler,
		userHandler,
	)
	routeSetup.Setup(app)

	return app
}
