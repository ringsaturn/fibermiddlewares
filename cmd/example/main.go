package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/oschwald/geoip2-golang"
	"github.com/ringsaturn/fibermiddlewares/xgeoip"
	"github.com/ringsaturn/fibermiddlewares/xhostname"
	"github.com/ringsaturn/fibermiddlewares/xresponsetime"
)

func GeoIPLocation(c *fiber.Ctx) error {
	val := c.Locals("X-GeoIP-City")

	city, ok := val.(*geoip2.City)
	if !ok {
		c = c.Status(500)
		return c.SendString("internal error")
	}
	return c.JSON(city.Location)
}

func main() {

	db, err := geoip2.Open("GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Use(xresponsetime.Handler)
	app.Use(xhostname.Handler)
	app.Use(xgeoip.New(db))

	app.Get("/", func(c *fiber.Ctx) error {
		time.Sleep(20 * time.Microsecond)
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/geoip/location", GeoIPLocation)
	_ = app.Listen("localhost:8999")
}
