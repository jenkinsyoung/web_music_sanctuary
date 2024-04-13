package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/database"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/hash"
	"os"
	"time"
)

type TokenClaims struct {
	jwt.MapClaims
	UserId int64 `json:"user_id"`
}

var SigningKey = os.Getenv("SECRET_KEY")

func GenerateToken(email, password string) (string, error) {
	user, err := database.DB.GetUserInfoByEmail(email)
	if err != nil {
		return "", err
	}
	if !hash.CheckPassword(password, user.Password) {
		return "", errors.New("wrong password")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.MapClaims{
			"ExpiresAt": time.Now().Add(time.Hour * 24 * 7).Unix(),
			"IssuedAt":  time.Now().Unix(),
		},
		user.Id,
	})

	token, err := t.SignedString([]byte(SigningKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(accessToken string) (int64, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(SigningKey), nil
	})
	if err != nil {
		return 0, err
	}

	fmt.Println(token, "\n")

	claims, ok := token.Claims.(*TokenClaims)
	fmt.Println(claims)
	//for id, i := range claims{
	//	fmt.Println(id, i)
	//}
	if !ok {
		return 0, errors.New("error in token claims")
	}
	return claims.UserId, nil
}
