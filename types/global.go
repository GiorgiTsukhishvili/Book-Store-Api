package types

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID uint
	Email  string
	jwt.RegisteredClaims
}
