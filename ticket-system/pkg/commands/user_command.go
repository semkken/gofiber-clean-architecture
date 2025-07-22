package commands

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"ticket-system/pkg/auth"
	"ticket-system/pkg/entities"
	"ticket-system/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type RegisterUserCommand struct {
	Username string
	Email    string
	Password string
}

type LoginUserCommand struct {
	Username string
	Password string
}

type UserCommandHandler interface {
	HandleRegisterUser(ctx context.Context, cmd RegisterUserCommand) error
	HandleLoginUser(ctx context.Context, cmd LoginUserCommand) (string, error)
}

type userCommandHandler struct {
	repo       repository.UserRepository
	jwtService *auth.JWTService
}

func NewUserCommandHandler(repo repository.UserRepository, jwtService *auth.JWTService) UserCommandHandler {
	return &userCommandHandler{repo: repo, jwtService: jwtService}
}

func (h *userCommandHandler) HandleRegisterUser(ctx context.Context, cmd RegisterUserCommand) error {
	idBytes := make([]byte, 16)
	_, err := rand.Read(idBytes)
	if err != nil {
		return err
	}
	id := hex.EncodeToString(idBytes)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(cmd.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entities.User{
		ID:       id,
		Username: cmd.Username,
		Email:    cmd.Email,
		Password: string(hashedPassword),
	}
	return h.repo.SaveUser(ctx, user)
}

func (h *userCommandHandler) HandleLoginUser(ctx context.Context, cmd LoginUserCommand) (string, error) {
	user, err := h.repo.FindUserByUsername(ctx, cmd.Username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cmd.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := h.jwtService.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
