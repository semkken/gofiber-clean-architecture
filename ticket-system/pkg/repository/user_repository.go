package repository

import (
	"context"
	"ticket-system/pkg/entities"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *entities.User) error
	FindUserByUsername(ctx context.Context, username string) (*entities.User, error)
	FindUserByID(ctx context.Context, id string) (*entities.User, error)
}
