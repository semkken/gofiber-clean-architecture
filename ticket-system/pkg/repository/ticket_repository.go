package repository

import (
	"context"
	"ticket-system/pkg/entities"
)

type TicketRepository interface {
	SaveTicket(ctx context.Context, ticket *entities.Ticket) error
	FindTicketByID(ctx context.Context, id string) (*entities.Ticket, error)
	FindAllTickets(ctx context.Context) ([]*entities.Ticket, error)
}
