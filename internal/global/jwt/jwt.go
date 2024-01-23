package jwt

import (
	"gin-rush-template/config"
	"github.com/golang-jwt/jwt"
	"time"
)

type Payload struct {
	UserId uint `json:"user_id"`
}

type Claims struct {
	Payload
	jwt.StandardClaims
}

// CreateToken 签发用户Token
func CreateToken(payload Payload) string {
	claims := Claims{
		Payload: payload,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(config.Get().JWT.AccessExpire)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := tokenClaims.SignedString([]byte(config.Get().JWT.AccessSecret))
	return token
}

// ParseToken 解析用户Token
func ParseToken(token string) (claims *Claims, ok bool) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{},
		func(token *jwt.Token) (any, error) {
			return []byte(config.Get().JWT.AccessSecret), nil
		},
	)
	if err != nil || tokenClaims == nil {
		return nil, false
	}
	if claims, ok = tokenClaims.Claims.(*Claims); !ok || !tokenClaims.Valid {
		return nil, false
	}
	return claims, true
}
