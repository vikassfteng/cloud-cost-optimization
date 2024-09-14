package main

import (
	"cloud-cost-optimizer/internal/optimization"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/optimize", func(c *gin.Context) {
		optimization.OptimizeCosts()
		c.JSON(200, gin.H{
			"message": "Optimization process completed.",
		})
	})

	r.Run() // Start the server
}
