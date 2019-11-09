package msjwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type TokenInterface interface {
	Generate(id string, isAdmin bool) (string, error)
	Validate(str string) (jwt.MapClaims, error)
}

type Token struct {
	Config *Config
}

func NewToken(config *Config) *Token {
	token := &Token{Config: config}

	return token
}

func (token *Token) Generate(id string, isAdmin bool) (string, error) {
	claims := struct {
		jwt.StandardClaims
		IsAdmin bool `json:"is_admin"`
	}{
		jwt.StandardClaims{
			Id:        uuid.New().String(),
			Audience:  token.Config.Audience,
			Subject:   id,
			ExpiresAt: time.Now().Add(5 * time.Hour).Unix(),
			Issuer:    token.Config.Issuer,
			IssuedAt:  time.Now().Unix(),
		},
		isAdmin,
	}

	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tk.SignedString([]byte(token.Config.Secret))

	return tokenString, err
}

func (token *Token) Validate(str string) (jwt.MapClaims, error) {
	tk, err := jwt.Parse(str, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(token.Config.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		return claims, nil
	}
	return nil, errors.New("could not validate")
}
