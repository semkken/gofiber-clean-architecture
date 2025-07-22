package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Logging() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		log.Printf("Started %s %s", c.Method(), c.Path())
		err := c.Next()
		log.Printf("Completed %s %s in %v", c.Method(), c.Path(), time.Since(start))
		return err
	}
}
