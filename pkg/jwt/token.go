package jwt_token

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func GenerateToken(user_id int64) (string, error) {

	claim := jwt.MapClaims{}
	claim["user_id"] = user_id
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claim["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	jwt_secret := os.Getenv("API_SECRET")
	tokenString, err := token.SignedString([]byte(jwt_secret))

	return tokenString, err
}
