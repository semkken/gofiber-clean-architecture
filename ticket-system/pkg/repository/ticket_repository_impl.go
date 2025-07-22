package repository

import (
	"context"
	"errors"
	"ticket-system/pkg/entities"
)

type ticketRepository struct {
	tickets map[string]*entities.Ticket
}

func NewTicketRepository() TicketRepository {
	return &ticketRepository{
		tickets: make(map[string]*entities.Ticket),
	}
}

func (r *ticketRepository) SaveTicket(ctx context.Context, ticket *entities.Ticket) error {
	r.tickets[ticket.ID] = ticket
	return nil
}

func (r *ticketRepository) FindTicketByID(ctx context.Context, id string) (*entities.Ticket, error) {
	ticket, exists := r.tickets[id]
	if !exists {
		return nil, errors.New("ticket not found")
	}
	return ticket, nil
}

func (r *ticketRepository) FindAllTickets(ctx context.Context) ([]*entities.Ticket, error) {
	tickets := make([]*entities.Ticket, 0, len(r.tickets))
	for _, ticket := range r.tickets {
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}
