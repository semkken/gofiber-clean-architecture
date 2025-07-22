package handlers

import (
	"context"
	"errors"
	"ticket-system/api/dto"
	"ticket-system/pkg/commands"
	"ticket-system/pkg/queries"

	"github.com/gofiber/fiber/v2"
)

type TicketHandler struct {
	cmdHandler   commands.TicketCommandHandler
	queryHandler queries.TicketQueryHandler
}

func NewTicketHandler(cmdHandler commands.TicketCommandHandler, queryHandler queries.TicketQueryHandler) *TicketHandler {
	return &TicketHandler{cmdHandler: cmdHandler, queryHandler: queryHandler}
}

func (h *TicketHandler) CreateTicket(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req dto.CreateTicketRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if req.Title == "" || req.Description == "" {
		return errors.New("invalid input")
	}
	cmd := commands.CreateTicketCommand{
		Title:       req.Title,
		Description: req.Description,
		CreatorID:   userID,
	}
	err := h.cmdHandler.HandleCreateTicket(context.Background(), cmd)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Ticket created"})
}

func (h *TicketHandler) AssignTicket(c *fiber.Ctx) error {
	var req dto.AssignTicketRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if req.AssigneeID == "" {
		return errors.New("invalid input")
	}
	cmd := commands.AssignTicketCommand{
		TicketID:   c.Params("id"),
		AssigneeID: req.AssigneeID,
	}
	err := h.cmdHandler.HandleAssignTicket(context.Background(), cmd)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "Ticket assigned"})
}

func (h *TicketHandler) GetTicket(c *fiber.Ctx) error {
	query := queries.GetTicketQuery{ID: c.Params("id")}
	ticket, err := h.queryHandler.HandleGetTicket(context.Background(), query)
	if err != nil {
		return err
	}
	response := dto.TicketResponseDTO{
		ID:          ticket.ID,
		Title:       ticket.Title,
		Description: ticket.Description,
		CreatorID:   ticket.CreatorID,
		AssigneeID:  ticket.AssigneeID,
	}
	return c.JSON(response)
}

func (h *TicketHandler) GetAllTickets(c *fiber.Ctx) error {
	tickets, err := h.queryHandler.HandleGetAllTickets(context.Background(), queries.GetAllTicketsQuery{})
	if err != nil {
		return err
	}
	response := make([]dto.TicketResponseDTO, len(tickets))
	for i, ticket := range tickets {
		response[i] = dto.TicketResponseDTO{
			ID:          ticket.ID,
			Title:       ticket.Title,
			Description: ticket.Description,
			CreatorID:   ticket.CreatorID,
			AssigneeID:  ticket.AssigneeID,
		}
	}
	return c.JSON(response)
}
