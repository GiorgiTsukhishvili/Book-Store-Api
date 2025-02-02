package utils

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/GiorgiTsukhishvili/BookShelf-Api/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWTInfo struct {
	Token                  string
	RefreshToken           string
	TokenExpiration        time.Time
	RefreshTokenExpiration time.Time
}

func GenerateJWTTokens(userID uint, email string) (*JWTInfo, error) {
	expiration, err := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRATION_DATE"))

	if err != nil {
		return nil, err
	}

	refreshExpiration, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION_DATE"))
	if err != nil {
		return nil, err
	}

	tokenExpiration := time.Now().Add(time.Duration(expiration) * time.Minute)
	refreshTokenExpiration := time.Now().Add(time.Duration(refreshExpiration) * time.Minute)

	claims := types.CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtRefreshToken, err := refreshToken.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	jwtInfo := JWTInfo{
		Token:                  jwtToken,
		RefreshToken:           jwtRefreshToken,
		TokenExpiration:        tokenExpiration,
		RefreshTokenExpiration: refreshTokenExpiration,
	}

	return &jwtInfo, nil
}

func ParseJwtToken(tokenString string, ctx *gin.Context, secret string) *types.CustomClaims {
	token, err := jwt.ParseWithClaims(tokenString, &types.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv(secret)), nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return nil
	}

	claims, ok := token.Claims.(*types.CustomClaims)

	if !ok || !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return nil
	}

	return claims
}
