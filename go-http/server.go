package main

import (
	"github.com/Ogury/profiling/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func main() {
	app := fiber.New()
	app.Use(pprof.New())

	service := &routes.MockService{}
	routes.AdRouter(app, service)
	app.Listen(":8080")
}
