package queries

import (
	"context"
	"ticket-system/pkg/entities"
	"ticket-system/pkg/repository"
)

type GetUserQuery struct {
	ID string
}

type UserQueryHandler interface {
	HandleGetUser(ctx context.Context, query GetUserQuery) (*entities.User, error)
}

type userQueryHandler struct {
	repo repository.UserRepository
}

func NewUserQueryHandler(repo repository.UserRepository) UserQueryHandler {
	return &userQueryHandler{repo: repo}
}

func (h *userQueryHandler) HandleGetUser(ctx context.Context, query GetUserQuery) (*entities.User, error) {
	return h.repo.FindUserByID(ctx, query.ID)
}
