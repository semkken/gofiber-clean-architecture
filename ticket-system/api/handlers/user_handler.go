package handlers

import (
	"context"
	"errors"
	"ticket-system/api/dto"
	"ticket-system/pkg/auth"
	"ticket-system/pkg/commands"
	"ticket-system/pkg/queries"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	cmdHandler   commands.UserCommandHandler
	queryHandler queries.UserQueryHandler
	jwtService   *auth.JWTService
}

func NewUserHandler(cmdHandler commands.UserCommandHandler, queryHandler queries.UserQueryHandler, jwtService *auth.JWTService) *UserHandler {
	return &UserHandler{cmdHandler: cmdHandler, queryHandler: queryHandler, jwtService: jwtService}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterUserRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return errors.New("invalid input")
	}
	cmd := commands.RegisterUserCommand{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	err := h.cmdHandler.HandleRegisterUser(context.Background(), cmd)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered"})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginUserRequestDTO
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if req.Username == "" || req.Password == "" {
		return errors.New("invalid input")
	}
	cmd := commands.LoginUserCommand{
		Username: req.Username,
		Password: req.Password,
	}
	token, err := h.cmdHandler.HandleLoginUser(context.Background(), cmd)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"token": token})
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	query := queries.GetUserQuery{ID: userID}
	user, err := h.queryHandler.HandleGetUser(context.Background(), query)
	if err != nil {
		return err
	}
	response := dto.UserResponseDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return c.JSON(response)
}
