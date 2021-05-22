package v1

import (
	"city-card-api/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reqLogin struct {
	Login    string `json:"email" bind:"required"`
	Password string `json:"password" bind:"required"`
}
type APIRes struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	// MessageR string `json:"messagerR"`
}

func (s *HttpV1) Login(c *gin.Context) {
	var body reqLogin
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error request login:", err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "check params",
			// "messageR": "Проверьте параметры",
		})
		return
	}
	user, tokens, err := s.auth.Login(body.Login, body.Password)
	if err != nil {
		c.JSON(404, gin.H{
			"error":   true,
			"message": "user_not_auth",
			// "messageR": "Логин и пароль не верны",
		})
		return
	}
	balance, err := s.pay.Balance(user.ID.Hex())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "error_get_balance",
			// "messageR": "Логин и пароль не верны",
		})
		return
	}
	c.JSON(200, gin.H{
		"error":   false,
		"message": "ok",
		"tokens":  tokens,
		"user":    user,
		"balance": balance,
	})
}

func (s *HttpV1) Refresh(c *gin.Context) {
	c.JSON(200, gin.H{
		"error":   false,
		"message": "ok",
	})
}

func (s *HttpV1) Register(c *gin.Context) {
	var body models.UserWithPassword
	if err := c.ShouldBindJSON(&body); err != nil {
		log.Println("Error request register:", err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "check params",
			// "messageR": "Не верно заполнены данные",
		})
		return
	}
	newUser, tokens, err := s.auth.Register(body)
	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "user not register, maybe login, email or telephone already used",
			// "messageR": "Пользователь не зарегистрирован, возможно логин, электронаня почта или телефон уже используются",
		})
		return
	}
	c.JSON(201, gin.H{
		"error":   false,
		"message": "ok",
		"tokens":  tokens,
		"user":    newUser,
		"balance": float32(0),
	})
}

func (h *HttpV1) Hello(c *gin.Context) {

	c.JSON(200, gin.H{
		"error":   false,
		"message": "ok",
	})
}
