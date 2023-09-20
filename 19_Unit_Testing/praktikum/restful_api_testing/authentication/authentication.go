package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var JwtSecret = []byte("secret-key") // Gantilah dengan kunci rahasia yang lebih aman

// CreateJWTToken membuat token JWT dengan klaim userID
func CreateJWTToken(userID uint) (string, error) {
	// Tentukan klaim token JWT
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku selama 1 hari (ganti sesuai kebutuhan Anda)
	}

	// Buat token JWT dengan klaim di atas dan tanda tangan menggunakan kunci rahasia
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyPassword(storedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}
