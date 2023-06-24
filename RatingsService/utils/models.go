package utils

import (
	"github.com/golang-jwt/jwt"
)

type Response struct {
	Message string `json:"Message"`
}

type Claims struct {
	Id   uint `json:"Id"`
	Role Role `json:"Role"`
	jwt.StandardClaims
}
