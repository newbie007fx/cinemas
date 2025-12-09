package authentication

import (
	"fmt"
	"time"

	"github.com/newbie007fx/cinemas/internal/module/users/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const (
	ACCESS_TOKEN  TokenType = "access_token"
	REFRESH_TOKEN TokenType = "refresh_token"
)

type TokenAuth struct {
	secret []byte
}

func New(secret string) *TokenAuth {
	return &TokenAuth{
		secret: []byte(secret),
	}
}

func (tg TokenAuth) GenerateAuthToken(user *models.User) (accessToken, refreshToken string, err error) {
	issuedAt := time.Now()
	claims := JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "github.com/newbie007fx/cinemas",
			IssuedAt: jwt.NewNumericDate(issuedAt),
		},
		Data: map[string]string{
			"id":       fmt.Sprint(user.ID),
			"name":     user.Name,
			"username": user.Username,
			"email":    user.Email,
		},
	}

	accessToken, err = tg.generateJwtToken(claims, ACCESS_TOKEN)
	if err != nil {
		return
	}
	refreshToken, err = tg.generateJwtToken(claims, REFRESH_TOKEN)

	return
}

func (tg TokenAuth) generateJwtToken(claims JwtCustomClaims, tokenType TokenType) (token string, err error) {
	expiredAt := time.Hour * 15
	if tokenType == REFRESH_TOKEN {
		expiredAt = time.Hour * 20
	}

	tokenId, err := uuid.NewV7()
	if err != nil {
		return
	}

	claims.ID = tokenId.String()
	claims.ExpiresAt = jwt.NewNumericDate(claims.IssuedAt.Add(expiredAt))
	claims.TokenType = tokenType

	newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = newtoken.SignedString(tg.secret)
	return
}

func (tg TokenAuth) ParseJWT(token string) (*JwtCustomClaims, error) {
	claims := &JwtCustomClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return (tg.secret), nil
	})

	return claims, err
}
