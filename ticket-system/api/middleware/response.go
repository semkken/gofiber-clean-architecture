package middleware

import (
	"ticket-system/api/dto"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ResponseMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		response := make(map[string]interface{})
		status := c.Response().StatusCode()

		if err != nil {
			response["success"] = false
			response["error"] = dto.ErrorResponseDTO{Error: err.Error()}.Error
			status = fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				status = e.Code
			}
		} else {
			response["success"] = true
			if body := c.Response().Body(); len(body) > 0 {
				response["data"] = c.JSON(body)
			} else {
				response["data"] = nil
			}
		}

		response["timestamp"] = time.Now().UTC().Format(time.RFC3339)
		response["duration_ms"] = time.Since(start).Milliseconds()

		c.Set("Content-Type", "application/json")
		c.Set("X-API-Version", "1.0")
		return c.Status(status).JSON(response)
	}
}
