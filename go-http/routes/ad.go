package routes

import (
	"github.com/Ogury/profiling/models"
	"github.com/gofiber/fiber/v2"
)

type AdService interface {
	HandleBidRequest(request *models.BidBodyRequest) (*models.BidResponse, error)
}

type MockService struct {
}

func (s *MockService) HandleBidRequest(request *models.BidBodyRequest) (*models.BidResponse, error) {
	return &models.BidResponse{}, nil
}

func AdRouter(app fiber.Router, service AdService) {
	app.Post("/ad", handle_ad(service))
}

func handle_ad(service AdService) fiber.Handler {
	return func(c *fiber.Ctx) error {
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
	}
}
