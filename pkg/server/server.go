package server

import (
	"fmt"
	"github.com/fardin01/fardin-payment-provider/pkg/payment"
	"github.com/gin-gonic/gin"
	"time"
)

func Start() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Nothing to see here.",
		})
	})

	r.GET("/readiness", func(c *gin.Context) {
		// Pretending that it takes 5 seconds for the app to start and to be ready to server requests.
		// In production, this should fail if code/app is not ready e.g. if it fails to connect to a database or a dependency.
		time.Sleep(5 * time.Second)
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/liveliness", func(c *gin.Context) {
		// In production, this should succeed as long as the webserver is able to respond.
		c.JSON(200, gin.H{
			"message": "Alive",
		})
	})

	r.POST("/rest/v1/payments/pay", func(c *gin.Context) {
		// payInvoices function in Antaeus calls this endpoint in a loop, which can very easily overload this server and
		// cause it to crash/suffer performance issues (not scalable). In production, This payment provider server should
		// offer a batch pay API, so Antaeus can pay n invoices with one API call.
		invoice, err := c.GetRawData()
		if err != nil {
			fmt.Println("Could not get request body: ", err)
			return
		}
		s := payment.Pay(invoice)
		c.JSON(s.StatusCode, gin.H{
			"result": s.Result,
		})
	})

	err := r.Run(":9000")

	if err != nil {
		fmt.Println("Could not start the server: ", err)
	}
}
