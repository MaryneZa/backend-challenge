package util

import (
	"fmt"
	"log"
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/MaryneZa/backend-challenge/internal/core/domain"
)

func CreateToken(user *domain.User, JWTSecret string) (string, error) {
	log.Println("secretKey:", JWTSecret, "|","user.ID:", user.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": user.ID.Hex(),
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})
	signedToken, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func VerifyToken(tokenString string, JWTSecret string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error, unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWTSecret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
            return "", errors.New("token has expired")
        }
        return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["userID"]
        if !ok {
            return "", errors.New("invalid user_id claim")
        }
        return userID, nil
	}
	return "", errors.New("invalid token")
}
