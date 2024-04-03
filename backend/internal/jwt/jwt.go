package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/database"
	"os"
	"time"
)

type TokenClaims struct {
	jwt.Claims
	UserId int `json:"user_id"`
}

var signingKey = os.Getenv("SECRET_KEY")

func GenerateToken(email, password string) (string, error) {
	userID, err := database.DB.GetUserID(email, password)
	fmt.Println("from jwt", password)
	if err != nil {
		return "", err
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.MapClaims{
			"ExpiresAt": time.Now().Add(time.Hour * 24 * 7).Unix(),
			"IssuedAt":  time.Now().Unix(),
		},
		userID,
	})

	token, err := t.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("error in token claims")
	}
	return claims.UserId, nil
}
