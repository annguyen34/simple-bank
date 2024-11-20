package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid secret key size, must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey: secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         payload.ID,
		"username":   payload.Username,
		"issued_at":  payload.IssuedAt.Unix(),
		"expired_at": payload.ExpiredAt.Unix(),
	})
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	if err != nil {
		return "", payload, err
	}
	return token, payload, nil
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte(maker.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &Payload{
		ID:        uuid.MustParse(claims["id"].(string)),
		Username:  claims["username"].(string),
		IssuedAt:  time.Unix(int64(claims["issued_at"].(float64)), 0),
		ExpiredAt: time.Unix(int64(claims["expired_at"].(float64)), 0),
	}, nil

}
