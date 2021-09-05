package xhostname

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {
	name, err := os.Hostname()
	if err != nil {
		name = "unknown"
	}
	c.Response().Header.Add("X-Host-Name", name)
	return c.Next()
}
