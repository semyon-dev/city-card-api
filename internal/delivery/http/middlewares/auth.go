package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *Middlewares) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		headerArray := strings.Split(header, " ")
		if len(headerArray) < 2 {
			c.AbortWithStatusJSON(400, gin.H{
				"errors":  true,
				"message": "check header Authorization",
			})
			return
		}
		if headerArray[0] != "Bearer" {
			c.AbortWithStatusJSON(400, gin.H{
				"errors":  true,
				"message": "check Bearer word in header Authorization",
			})
			return
		}
		token := headerArray[1]
		user, err := m.Decode(token)
		if err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"errors":  true,
				"message": "error decode token",
			})
			return
		}
		c.Set("user_id", user.UserID)
		c.Set("user_role", user.Role)
		c.Next()
	}
}
