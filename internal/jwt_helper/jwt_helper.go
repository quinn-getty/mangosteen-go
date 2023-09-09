package jwt_helper

import (
	"crypto/rand"
	"io"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateJWT(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
	})

	key, err := getHMACKey()
	if err != nil {
		return "", err
	}

	return token.SignedString(key)
}

func getHMACKey() ([]byte, error) {
	keyPath := viper.GetString("jwt.hmac.key_path")
	return os.ReadFile(keyPath)
}

func GenerateHMACKey() ([]byte, error) {
	key := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, err
	}
	return key, nil
}
