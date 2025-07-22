package main

import (
	"ticket-system/api/handlers"
	"ticket-system/api/routes"
	"ticket-system/pkg/auth"
	"ticket-system/pkg/commands"
	"ticket-system/pkg/queries"
	"ticket-system/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize repositories
	userRepo := repository.NewUserRepository()
	ticketRepo := repository.NewTicketRepository()

	// Initialize auth service
	jwtService := auth.NewJWTService("your-secret-key") // Replace with secure key

	// Initialize command and query handlers
	userCmdHandler := commands.NewUserCommandHandler(userRepo, jwtService)
	userQueryHandler := queries.NewUserQueryHandler(userRepo)

	ticketCmdHandler := commands.NewTicketCommandHandler(ticketRepo, userRepo)
	ticketQueryHandler := queries.NewTicketQueryHandler(ticketRepo)

	// Initialize HTTP handlers
	userHandler := handlers.NewUserHandler(userCmdHandler, userQueryHandler, jwtService)
	ticketHandler := handlers.NewTicketHandler(ticketCmdHandler, ticketQueryHandler)

	// Setup routes
	routes.SetupRoutes(app, userHandler, ticketHandler)

	// Start server
	app.Listen(":3000")
}
