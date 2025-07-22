package routes

import (
	"ticket-system/api/handlers"
	"ticket-system/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *handlers.UserHandler, ticketHandler *handlers.TicketHandler) {
	app.Use(middleware.Logging())
	app.Use(middleware.ResponseMiddleware())

	// Public routes
	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	// Protected routes
	api := app.Group("/api", middleware.JWTAuth(userHandler.JWTService))
	api.Get("/profile", userHandler.Profile)

	// Ticket routes
	tickets := api.Group("/tickets")
	tickets.Post("/", ticketHandler.CreateTicket)
	tickets.Post("/:id/assign", ticketHandler.AssignTicket)
	tickets.Get("/:id", ticketHandler.GetTicket)
	tickets.Get("/", ticketHandler.GetAllTickets)
}
