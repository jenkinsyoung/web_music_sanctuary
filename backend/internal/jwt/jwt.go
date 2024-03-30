package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/hash"
	"github.com/jenkinsyoung/web_music_sanctuary/internal/user"
	"time"
)

type TokenClaims struct {
	jwt.Claims
	UserId int `json:"user_id"`
}

func GenerateToken(email, password string) (string, error) {
	hashedPassword := hash.PasswordHash(password)
	userID, err := user.GetUserID(email, hashedPassword)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.MapClaims{
			"ExpiresAt": time.Now().Add(time.Hour * 24 * 7).Unix(),
			"IssuedAt":  time.Now().Unix(),
		},
		userID,
	})

	return token.SigningString([]byte(signingKey))
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		nil, errors.New("invalid token claims")
	}

	return claims.UserId, nil
}
