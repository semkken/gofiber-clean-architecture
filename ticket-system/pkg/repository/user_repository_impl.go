package repository

import (
	"context"
	"errors"
	"ticket-system/pkg/entities"
)

type userRepository struct {
	users map[string]*entities.User
}

func NewUserRepository() UserRepository {
	return &userRepository{
		users: make(map[string]*entities.User),
	}
}

func (r *userRepository) SaveUser(ctx context.Context, user *entities.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *userRepository) FindUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *userRepository) FindUserByID(ctx context.Context, id string) (*entities.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
