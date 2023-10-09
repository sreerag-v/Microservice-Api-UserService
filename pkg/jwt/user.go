package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateRefreshToken implements interfaces.JWTUsecase
func GenerateJWT(id uint) (string, error) {

	expireTime := time.Now().Add(60 * time.Minute).Unix()

	// create token with expire time and claims id as user id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expireTime,
		Id:        fmt.Sprint(id),// converting into string 
	})

	// conver the token into signed string
	tokenString, err := token.SignedString([]byte("123456789"))

	if err != nil {
		return "", err
	}
	// refresh token add next time
	return  tokenString, nil
}


func ValidateToken(tokenString string) (jwt.StandardClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("123456789"), nil
		},
	)
	if err != nil || !token.Valid {
		return jwt.StandardClaims{}, errors.New("not valid token")
	}

	// then parse the token to claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return jwt.StandardClaims{}, errors.New("can't parse the claims")
	}

	return *claims, nil
}