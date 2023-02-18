package util

import (
	"encoding/json"
	"errors"

	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type MetaToken struct {
	ID            string
	Email         string
	ExpiredAt     time.Time
	Authorization bool
	Role int
}

type AccessToken struct {
	Claims MetaToken
}

func Sign(Data map[string]interface{}, SecretPublicKeyEnvName string, ExpiredAt time.Duration) (string, error) {
	expiredAt := time.Now().Add(time.Duration(time.Minute) * ExpiredAt).Unix()
	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(jwtSecretKey))
	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}
	return accessToken, nil
}

func VerifyTokenHeader(ctx *gin.Context, SecrePublicKeyEnvName string) (*jwt.Token, error) {
	splitToken := strings.SplitAfter(ctx.GetHeader("Authorization"), "Bearer")
	if len(splitToken) < 2 {
		err := errors.New("error: Authorization not found")
		logrus.Error(err.Error())
		return nil, err
	}
	accessToken := splitToken[1]
	jwtSecretKey := GodotEnv(SecrePublicKeyEnvName)

	token, err := jwt.Parse(strings.Trim(accessToken, " "), func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func VerifyToken(accessToken, SecrePublicKeyEnvName string) (*jwt.Token, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(GodotEnv(SecrePublicKeyEnvName)), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken {
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify), &token)

	return token
}

func HashPassword(password string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return string(result)
}

func ComparePassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}