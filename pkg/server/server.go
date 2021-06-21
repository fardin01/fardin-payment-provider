package server

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Start()  {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Nothing to see here.",
		})
	})

	r.GET("/readiness", func(c *gin.Context) {
		// Pretending that it takes 10 seconds for the app to start and to be ready to server requests.
		// In production, this should fail if code/app is not ready e.g. if it fails to connect to a database or dependency.
		time.Sleep(10 * time.Second)
		c.JSON(200, gin.H{
			"message": "Ready",
		})
	})

	r.GET("/liveliness", func(c *gin.Context) {
		// In production, this should succeed as long as the webserver is able to respond.
		c.JSON(200, gin.H{
			"message": "Alive",
		})
	})

	r.POST("/rest/v1/payments/", func(c *gin.Context) {
		c.JSON(201, gin.H{
			"result": true,
		})
	})

	r.Run(":9000")
}
