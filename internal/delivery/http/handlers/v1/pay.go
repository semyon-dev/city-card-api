package v1

import (
	_ "city-card-api/internal/repository"
	_ "city-card-api/internal/services"
	"log"
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

func (v1 HttpV1) RequestPay(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	token, err := v1.pay.RequestPay(userID)
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
		"token":   token,
	})
}

type reqApprovePay struct {
	PayToken string  `json:"payToken" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
}

func (v1 HttpV1) ApprovePay(c *gin.Context) {
	var body reqApprovePay
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "check_params",
		})
		return
	}
	toUserID := c.MustGet("user_id").(string)
	err := v1.pay.SubmitPay(toUserID, body.PayToken, body.Amount)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error":   true,
			"message": "no_money",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "ok",
	})
}
