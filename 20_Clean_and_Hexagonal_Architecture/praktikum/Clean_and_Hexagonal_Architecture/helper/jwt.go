package helper

import (
	"Clean_and_Hexagonal_Architecture/constants"
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(10 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
