package utils

import (
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func HashedPassword(password string) (string, error) {
	// 生成密码的哈希值
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func GenerateJWT(username string) (string, error) {
	// 生成jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return "Bearer " + signedToken, err
}

func CheckPassword(input string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(input))
	return err == nil
}
