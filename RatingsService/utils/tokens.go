package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
)

// todo ovo preko env varijable namestiti
var JwtKey = []byte("z7031Q8Qy9zVO-T2o7lsFIZSrd05hH0PaeaWIBvLh9s")

func ParseTokenStr(tokenStr string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("error while parsing jwt str")
		}
		return JwtKey, nil
	})

	return token, err
}

func GetClaimsFromHeader(r *http.Request) (*Claims, error) {
	bearer := r.Header["Authorization"]
	if bearer == nil {
		return nil, errors.New("no token")
	}

	tokenStr := strings.Split(bearer[0], " ")[1]
	token, err := ParseTokenStr(tokenStr)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	claims := token.Claims.(jwt.MapClaims)
	id_str := fmt.Sprintf("%v", claims["Id"])
	role_str := fmt.Sprintf("%v", claims["Role"])
	id, err := strconv.ParseUint(id_str, 10, 32)
	if err != nil {
		panic(err)
	}
	claims_basic := Claims{Id: uint(id), Role: Role(role_str)}

	return &claims_basic, nil
}

type Role string

const (
	HOST  Role = "HOST"
	GUEST Role = "GUEST"
)
