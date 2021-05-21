package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	SYSTEM_ROLE_USER    = "user"
	SYSTEM_ROLE_WORKER  = "worker"
	SYSTEM_ROLE_MANAGER = "manager"
	SYSTEM_ROLE_ADMIN   = "admin"
)

type UserProfile struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" binding:"required"`
	Surname   string             `json:"surname" binding:"required"`
	MName     string             `json:"mname" binding:"required"`
	Email     string             `json:"email" binding:"required"`
	Telephone string             `json:"telephone" binding:"required"`
	Role      string             `json:"role"`
}
type UserWithPassword struct {
	UserProfile `bson:",inline"`
	Password    string `json:"password" binding:"required"`
}

type UserJWT struct {
	jwt.StandardClaims
	UserID string
	Role   string
}
type Tokens struct {
	AccessToken string `json:"accessToken"`
	RefresToken string `json:"refreshToken"`
}
