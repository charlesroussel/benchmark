package routes

import (
	"sync"

	"github.com/Ogury/profiling/connectors"
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
	CallDynamoDB(request *models.BidBodyRequest) (*models.BidResponse, error)
	CallDax(request *models.BidBodyRequest) (*models.BidResponse, error)
}

type MockService struct {
}

func (s *MockService) HandleBidRequest(request *models.BidBodyRequest) (*models.BidResponse, error) {
	return &models.BidResponse{}, nil
}

func (s *MockService) CallDynamoDB(request *models.BidBodyRequest) (*models.BidResponse, error) {
	id, err := connectors.DynamoClient.Get("table_name", "id", false)
	if err != nil {
		return nil, err
	}
	return &models.BidResponse{
		ID: id,
	}, nil
}

func (s *MockService) CallDax(request *models.BidBodyRequest) (*models.BidResponse, error) {
	id, err := connectors.DynamoClient.Get("table_name", "id", true)
	if err != nil {
		return nil, err
	}
	return &models.BidResponse{
		ID: id,
	}, nil
}

func AdRouter(app fiber.Router, service AdService) {
	app.Post("/pool", handle_pool_ad(service))
	app.Post("/ad", handle_ad(service))
	app.Post("/dynamo", handle_dynamo(service))
	app.Post("/dax", handle_dax(service))
}

func handle_pool_ad(service AdService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := BidRequestPool.Get().(*models.BidBodyRequest)
		defer BidRequestPool.Put(request)
		if err := c.BodyParser(request); err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		return handle_result(c)(service.HandleBidRequest(request))
	}
}

func handle_ad(service AdService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := &models.BidBodyRequest{}
		if err := c.BodyParser(request); err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		return handle_result(c)(service.HandleBidRequest(request))
	}
}

func handle_dynamo(service AdService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := BidRequestPool.Get().(*models.BidBodyRequest)
		defer BidRequestPool.Put(request)
		if err := c.BodyParser(request); err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		return handle_result(c)(service.CallDynamoDB(request))
	}
}

func handle_dax(service AdService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		request := BidRequestPool.Get().(*models.BidBodyRequest)
		defer BidRequestPool.Put(request)
		if err := c.BodyParser(request); err != nil {
			_ = c.JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}

		return handle_result(c)(service.CallDax(request))
	}
}

type ResultHandler = func(result *models.BidResponse, err error) error

func handle_result(c *fiber.Ctx) ResultHandler {
	return func(result *models.BidResponse, err error) error {
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
