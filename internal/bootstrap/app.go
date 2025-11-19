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
	transactionRepo := repository.NewTransactionRepository(config.DB)
	
	authService := services.NewAuthService(userRepo, tukangRepo)
	userService := services.NewUserService(userRepo)
	tukangService := services.NewTukangService(tukangRepo)
	transactionService := services.NewTransactionService(transactionRepo, tukangRepo)
	
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(*userService)
	tukangHandler := handler.NewTukangHandler(*tukangService)
	orderHandler := handler.NewOrderHandler(transactionService)

	// Register routes
	routeSetup := routes.NewRouteSetup(
		authHandler,
		userHandler,
		tukangHandler,
		orderHandler,
	)
	routeSetup.Setup(app)

	return app
}
