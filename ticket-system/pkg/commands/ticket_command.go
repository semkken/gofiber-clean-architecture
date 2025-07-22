package commands

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"ticket-system/pkg/entities"
	"ticket-system/pkg/repository"
)

type CreateTicketCommand struct {
	Title       string
	Description string
	CreatorID   string
}

type AssignTicketCommand struct {
	TicketID   string
	AssigneeID string
}

type AttachFileCommand struct {
	TicketID string
	FilePath string
}

type TicketCommandHandler interface {
	HandleCreateTicket(ctx context.Context, cmd CreateTicketCommand) error
	HandleAssignTicket(ctx context.Context, cmd AssignTicketCommand) error
}

type ticketCommandHandler struct {
	ticketRepo repository.TicketRepository
	userRepo   repository.UserRepository
}

func NewTicketCommandHandler(ticketRepo repository.TicketRepository, userRepo repository.UserRepository) TicketCommandHandler {
	return &ticketCommandHandler{ticketRepo: ticketRepo, userRepo: userRepo}
}

func (h *ticketCommandHandler) HandleCreateTicket(ctx context.Context, cmd CreateTicketCommand) error {
	idBytes := make([]byte, 16)
	_, err := rand.Read(idBytes)
	if err != nil {
		return err
	}
	id := hex.EncodeToString(idBytes)

	ticket := &entities.Ticket{
		ID:          id,
		Title:       cmd.Title,
		Description: cmd.Description,
		CreatorID:   cmd.CreatorID,
		AssigneeID:  "",
	}
	return h.ticketRepo.SaveTicket(ctx, ticket)
}

func (h *ticketCommandHandler) HandleAssignTicket(ctx context.Context, cmd AssignTicketCommand) error {
	ticket, err := h.ticketRepo.FindTicketByID(ctx, cmd.TicketID)
	if err != nil {
		return err
	}
	_, err = h.userRepo.FindUserByID(ctx, cmd.AssigneeID)
	if err != nil {
		return errors.New("assignee not found")
	}
	ticket.AssigneeID = cmd.AssigneeID
	return h.ticketRepo.SaveTicket(ctx, ticket)
}
