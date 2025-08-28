package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/leandrowiemesfilho/login-api/internal/configs"
	"github.com/leandrowiemesfilho/login-api/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(foundPwd, pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(foundPwd), []byte(pwd))

	return err
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func CreateJWT(user *models.User) string {
	appConfig := configs.AppConfig
	secret := appConfig.JWTSecret
	claims := &jwt.MapClaims{
		"expiresAt":     appConfig.JWTExpiration,
		"accountNumber": user.Id.Hex(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	return signedString
}
