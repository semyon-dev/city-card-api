package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *HttpV1) Hello(c *gin.Context) {

	c.JSON(200, gin.H{
		"error":   false,
		"message": "ok",
	})
}
