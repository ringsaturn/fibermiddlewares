package xresponsetime

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func Handler(c *fiber.Ctx) error {
	start := time.Now()
	defer func() {
		c.Response().Header.Add("x-Response-Time", strconv.Itoa(int(time.Since(start).Microseconds())))
	}()
	return c.Next()
}
