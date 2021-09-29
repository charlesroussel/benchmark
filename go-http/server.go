package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/Ogury/profiling/models"
	"github.com/Ogury/profiling/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

func main() {
	fmt.Printf("Found number of cpus: %d\n", runtime.NumCPU())
	env_core_count, ok := os.LookupEnv("CORE_COUNT")
	if ok {
		core_uses, err := strconv.Atoi(env_core_count)
		if err == nil && core_uses > 0 {
			fmt.Printf("Using cores: %d\n", core_uses)
			runtime.GOMAXPROCS(core_uses)
		} else {
			fmt.Printf("Unable to parse: %s to a number\n", env_core_count)
		}
	}

	app := fiber.New()
	app.Use(pprof.New())

	service := &routes.MockService{}
	routes.AdRouter(app, service)
	app.Post("/bypass", func(c *fiber.Ctx) error {
		var adr models.BidBodyRequest
		if err := c.BodyParser(&adr); err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result, err := service.HandleBidRequest(&adr)

		if err != nil {
			return c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		return c.JSON(&fiber.Map{
			"success": true,
			"result":  result,
		})
	})
	app.Listen(":8080")
}
