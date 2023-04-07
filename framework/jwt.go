package framework

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/scarlet0725/prism-api/model"
)

type JWTHandler interface {
	CreateAccessToken(*model.User) (string, error)
	CreateStateToken(*model.User) (string, error)
	ValidateToken(token string) (*JWTPayload, error)
}

type JWTCustomClaims struct {
	ID     int    `json:"id"`
	UserID string `json:"userID"`
	jwt.RegisteredClaims
}

type JWTPayload struct {
	ID     int    `json:"id"`
	UserID string `json:"userID"`
	Type   string `json:"type"`
}

type JWTHandlerImpl struct {
	signingKey []byte
}

func NewJWTHandler(signingKey string) JWTHandler {
	return &JWTHandlerImpl{
		signingKey: []byte(signingKey),
	}
}

func (j *JWTHandlerImpl) CreateAccessToken(user *model.User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTCustomClaims{
		ID:     user.ID,
		UserID: user.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "prism-api",
			Audience: []string{
				"access_token",
			},
			Subject:   user.UserID,
			ExpiresAt: jwt.NewNumericDate(jwt.TimeFunc().Add(24 * 1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(jwt.TimeFunc()),
		}},
	)
	return t.SignedString(j.signingKey)
}

func (j *JWTHandlerImpl) CreateStateToken(user *model.User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTCustomClaims{
		ID:     user.ID,
		UserID: user.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "prism-api",
			Audience: []string{
				"state",
			},
			Subject:   user.UserID,
			ExpiresAt: jwt.NewNumericDate(jwt.TimeFunc().Add(24 * 1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(jwt.TimeFunc()),
		}},
	)

	return t.SignedString(j.signingKey)
}

func (j *JWTHandlerImpl) ValidateToken(tokenString string) (*JWTPayload, error) {
	t := &JWTCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, t, func(token *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if len(t.RegisteredClaims.Audience) == 0 {
		return nil, fmt.Errorf("invalid token")
	}

	p := &JWTPayload{}

	switch t.RegisteredClaims.Audience[0] {
	case "access_token":
		p.Type = "access_token"
	case "state":
		p.Type = "state"
	default:
		return nil, fmt.Errorf("invalid token")
	}

	return p, nil

}
