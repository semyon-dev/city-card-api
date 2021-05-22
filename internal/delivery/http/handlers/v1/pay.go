package v1

import (
	_ "city-card-api/internal/repository"
	_ "city-card-api/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (v1 HttpV1) Balance(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	balance, err := v1.pay.Balance(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "ok",
		"balance": balance,
	})
}

type reqAddMoney struct {
	Money float64 `json:"money" binding:"required"`
}

func (v1 HttpV1) AddMoney(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	var body reqAddMoney
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "check_params",
		})
		return
	}
	balance, err := v1.pay.AddMoney(userID, body.Money)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "ok",
		"balance": balance,
	})
}
