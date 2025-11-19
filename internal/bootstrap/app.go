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
	tukangRepo := repository.NewTukangRepository(config.DB)
	
	authService := services.NewAuthService(userRepo, tukangRepo)
	userService := services.NewUserService(userRepo)
	tukangService := services.NewTukangService(tukangRepo)
	
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(*userService)
	tukangHandler := handler.NewTukangHandler(*tukangService)

	// Register routes
	routeSetup := routes.NewRouteSetup(
		authHandler,
		userHandler,
		tukangHandler,
	)
	routeSetup.Setup(app)

	return app
}
