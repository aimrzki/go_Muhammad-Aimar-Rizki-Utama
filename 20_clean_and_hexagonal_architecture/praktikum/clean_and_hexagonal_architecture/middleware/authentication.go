package middleware

import (
	"clean_and_hexagonal_architecture/internal/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type JwtCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(user *model.User, secretKey string) (string, error) {
	claims := JwtCustomClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractTokenFromHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrInvalidToken
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return "", ErrInvalidToken
	}

	return tokenParts[1], nil
}

func JWTAuthMiddleware(secretKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString, err := ExtractTokenFromHeader(c)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}

			token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					return echo.NewHTTPError(http.StatusUnauthorized, ErrInvalidToken.Error())
				}
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}

			claims, ok := token.Claims.(*JwtCustomClaims)
			if !ok || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, ErrInvalidToken.Error())
			}

			c.Set("user_id", claims.UserID)

			return next(c)
		}
	}
}
