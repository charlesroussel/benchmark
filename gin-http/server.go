package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Ogury/profiling/models"
	"github.com/gin-gonic/gin"
)

func HandleBidRequest(request *models.BidBodyRequest) (*models.BidResponse, error) {
	return &models.BidResponse{}, nil
}

func HandleAd(c *gin.Context) {
	var request models.BidBodyRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	out, err := HandleBidRequest(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  out,
	})
}

var BidRequestPool = sync.Pool{
	New: func() interface{} {
		return new(models.BidBodyRequest)
	},
}

func HandleAdPool(c *gin.Context) {
	var request = BidRequestPool.Get().(*models.BidBodyRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	out, err := HandleBidRequest(request)
	BidRequestPool.Put(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  out,
	})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/ad", HandleAd)
	r.POST("/ad_pool", HandleAdPool)
	fmt.Println("Listening on 0.0.0.0:8080")
	r.Run("0.0.0.0:8080")
}
