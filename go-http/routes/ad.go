package routes

import (
	"sync"

	"github.com/Ogury/profiling/models"
	"github.com/gofiber/fiber/v2"
)

var BidRequestPool = sync.Pool{
	New: func() interface{} {
		return new(models.BidBodyRequest)
	},
}

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
		request := BidRequestPool.Get().(*models.BidBodyRequest)
		defer BidRequestPool.Put(request)
		if err := c.BodyParser(request); err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		result, err := service.HandleBidRequest(request)

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
