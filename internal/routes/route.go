package routes

import (
	"2Kang/internal/handler"
	"2Kang/internal/middleware"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type RouteSetup struct {
	authHandler       *handler.AuthHandler
	userHandler      *handler.UserHandler
	tukangHandler *handler.TukangHandler
	orderHandler *handler.OrderHandler
	tukangHomeHandler *handler.TukangHomeHandler
	tukangOrderHandler *handler.TukangOrderHandler
	tukangProfileHandler *handler.TukangProfileHandler
}

func NewRouteSetup(
	authHandler *handler.AuthHandler,
	userHandler *handler.UserHandler,
	tukangHandler *handler.TukangHandler,
	orderHandler *handler.OrderHandler,
	tukangHomeHandler *handler.TukangHomeHandler,
	tukangOrderHandler *handler.TukangOrderHandler,
	tukangProfileHandler *handler.TukangProfileHandler,
) *RouteSetup {
	return &RouteSetup{
		authHandler: authHandler,
		userHandler:  userHandler,
		tukangHandler: tukangHandler,
		orderHandler: orderHandler,
		tukangHomeHandler: tukangHomeHandler,
		tukangOrderHandler: tukangOrderHandler,
		tukangProfileHandler: tukangProfileHandler,
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

	api.Post("/register", rs.authHandler.Register)
	api.Post("/login", rs.authHandler.Login)

	protected := api.Group("/", middleware.JWTProtected)
	{
		// USER
		protected.Get("/user/profile", rs.userHandler.GetProfile)
		protected.Patch("/user/name", rs.userHandler.UpdateName)

		// TUKANG
		protected.Get("/tukang", rs.tukangHandler.GetTukangList)
		protected.Get("/tukang/:user_id", rs.tukangHandler.GetTukangDetail)

		protected.Patch("/tukang/category", rs.tukangHandler.UpdateCategory)
		protected.Patch("/tukang/bio", rs.tukangHandler.UpdateBio)
		protected.Patch("/tukang/services", rs.tukangHandler.UpdateServices)

		protected.Get("/tukang/home", rs.tukangHomeHandler.GetHome)

		// inside protected group
		protected.Post("/orders", rs.orderHandler.CreateOrder)
		protected.Get("/orders", rs.orderHandler.GetMyOrders)
		protected.Get("/orders/:id", rs.orderHandler.GetOrderDetail)

		// tukang endpoints
		protected.Get("/tukang/orders", rs.orderHandler.GetOrdersForTukang)
		protected.Patch("/tukang/orders/status", rs.orderHandler.UpdateOrderStatus)

		// TUKANG ORDER
		protected.Post("/tukang/order/:id/accept", rs.tukangOrderHandler.Accept)
		protected.Post("/tukang/order/:id/start", rs.tukangOrderHandler.Start)
		protected.Post("/tukang/order/:id/finish", rs.tukangOrderHandler.Finish)

		protected.Get("/tukang/profile", rs.tukangProfileHandler.GetProfile)
		protected.Patch("/tukang/profile", rs.tukangProfileHandler.UpdateProfile)

	}

}