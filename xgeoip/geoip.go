package xgeoip

import (
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/oschwald/geoip2-golang"
)

func New(reader *geoip2.Reader) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ip net.IP
		forwardedIP := net.ParseIP(c.Get("X-Forwarded-For"))
		if forwardedIP != nil {
			ip = forwardedIP
		} else {
			ip = c.Context().RemoteIP()
		}
		city, _ := reader.City(ip)
		c.Locals("X-GeoIP-City", city)

		return c.Next()
	}
}
