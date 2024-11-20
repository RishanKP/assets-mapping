package jwt

import (
	"fmt"
	"time"

	"asset-mapping/library/config"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserId   string
	Username string
}

var jwt_key = []byte(config.JWT_SECRET)

func CreateToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"username": claims.Username,
		"userId":   claims.UserId,
	})

	tokenString, err := token.SignedString(jwt_key)
	if err != nil {
		fmt.Println("Error sigining token: ", err)
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwt_key, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return Claims{}, err
	}

	tokenClaims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		fmt.Println("Error parsing token:", err)
		return Claims{}, err
	}

	claims := Claims{
		UserId:   tokenClaims["userId"].(string),
		Username: tokenClaims["username"].(string),
	}

	return claims, nil
}
