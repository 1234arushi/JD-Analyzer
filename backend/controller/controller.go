package controller

import (
	"github.com/1234arushi/JD-Analyzer/process"
	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context) {
	var (
		err error
		req process.JDRequest
	)
	if err = (&req).Validate(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err = (&req).Process(c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return

	}
	//c.JSON() -> sends JSON response back to the client
	c.JSON(200, gin.H{
		"message": "Job processed successfully",
	})

}
