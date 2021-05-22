package services

import (
	"city-card-api/internal/models"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var key = []byte("my-32-character-ultra-secure-and-ultra-long-secret")

func (auths *authService) Decode(token string) (*models.UserJWT, error) {
	myToken, err := jwt.ParseWithClaims(token, &models.UserJWT{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Println(err)
		return &models.UserJWT{}, err
	}
	if claims, ok := myToken.Claims.(*models.UserJWT); ok {
		return claims, nil
	}
	return &models.UserJWT{}, errors.New("non UserJWT type")
}

func (auths *authService) encode(timeHours int, userID string, role string) (string, error) {
	claims := models.UserJWT{}
	claims.UserID = userID
	claims.Role = role
	duration := time.Duration(int64(time.Hour) * int64(timeHours))
	claims.ExpiresAt = time.Now().Add(duration).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, err := token.SignedString(key)
	// log.Println("Error secret key:", err)
	return tokenString, err
}
func (auths *authService) encodeToAccessToken(userID string, role string) (string, error) {
	return auths.encode(24, userID, role)
}

func (auths *authService) encodeToRefreshToken(userID string, role string) (string, error) {
	return auths.encode(24*7, userID, role)
}

func (auths *authService) encodeUser(userID string, role string) (models.Tokens, error) {
	tokens := models.Tokens{}
	accessToken, err := auths.encodeToAccessToken(userID, role)
	if err != nil {
		log.Println(err)
		return tokens, err
	}
	refreshToken, err := auths.encodeToRefreshToken(userID, role)
	if err != nil {
		log.Println(err)
		return tokens, err
	}
	tokens = models.Tokens{
		AccessToken: accessToken,
		RefresToken: refreshToken,
	}
	return tokens, nil
}
func (auths *authService) Login(login, pass string) (models.UserProfile, models.Tokens, error) {
	var err error
	user, err := auths.db.ReadByLoginAndPass(login, pass)
	tokens := models.Tokens{}
	if err != nil {
		log.Println(err)
		return models.UserProfile{}, tokens, err
	}
	tokens, err = auths.encodeUser(user.ID.Hex(), user.Role)
	if err != nil {
		log.Println(err)
		return models.UserProfile{}, tokens, err
	}
	return user, tokens, nil
}

func (auths *authService) Register(user models.UserWithPassword) (models.UserProfile, models.Tokens, error) {
	newUser, err := auths.db.CreateUser(user)
	tokens := models.Tokens{}
	if err != nil {
		log.Println("Error create user in db:", err)
		return models.UserProfile{}, tokens, err
	}
	tokens, err = auths.encodeUser(user.ID.Hex(), user.Role)
	if err != nil {
		log.Println("Error create token:", err)
		return models.UserProfile{}, tokens, err
	}
	_, err = auths.dbPay.CreateCard(newUser.ID)
	if err != nil {
		log.Println("Error create token:", err)
		return models.UserProfile{}, tokens, err
	}
	return newUser, tokens, nil
}
