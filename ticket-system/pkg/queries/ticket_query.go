package queries

import (
	"context"
	"ticket-system/pkg/entities"
	"ticket-system/pkg/repository"
)

type GetTicketQuery struct {
	ID string
}

type GetAllTicketsQuery struct{}

type TicketQueryHandler interface {
	HandleGetTicket(ctx context.Context, query GetTicketQuery) (*entities.Ticket, error)
	HandleGetAllTickets(ctx context.Context, query GetAllTicketsQuery) ([]*entities.Ticket, error)
}

type ticketQueryHandler struct {
	repo repository.TicketRepository
}

func NewTicketQueryHandler(repo repository.TicketRepository) TicketQueryHandler {
	return &ticketQueryHandler{repo: repo}
}

func (h *ticketQueryHandler) HandleGetTicket(ctx context.Context, query GetTicketQuery) (*entities.Ticket, error) {
	return h.repo.FindTicketByID(ctx, query.ID)
}

func (h *ticketQueryHandler) HandleGetAllTickets(ctx context.Context, query GetAllTicketsQuery) ([]*entities.Ticket, error) {
	return h.repo.FindAllTickets(ctx)
}
