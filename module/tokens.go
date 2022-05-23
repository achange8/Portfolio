package module

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func CreateAccToken(userId string) (string, error) {
	envERR := godotenv.Load("db.env")
	if envERR != nil {
		fmt.Println("could not load env file !")
		os.Exit(1)
	}
	key := os.Getenv("key")
	claims := jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
	}
	rawtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawtoken.SignedString([]byte(key))
	if err != nil {
		return "failed signed env", err
	}
	return token, nil
}

func CreateRefreshToken(userId string) (string, error) {
	envERR := godotenv.Load("db.env")
	if envERR != nil {
		fmt.Println("could not load env file !")
		os.Exit(1)
	}
	key := os.Getenv("key2")
	claims := jwt.StandardClaims{
		Id:        userId,
		ExpiresAt: time.Now().Add(24 * 30 * time.Hour).Unix(),
	}
	rawtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := rawtoken.SignedString([]byte(key))
	if err != nil {
		return "failed signed env", err
	}
	return token, nil
}
