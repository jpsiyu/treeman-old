package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jpsiyu/treeman/server/conf"
)

type AuthUser struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

const expirationTime = 5 * time.Minute

func GenTokenStr(user *AuthUser) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["iat"] = time.Now().Unix()
	claims["data"] = *user
	claims["exp"] = time.Now().Add(expirationTime).Unix()
	tokenStr, err := token.SignedString([]byte(conf.AuthKey))
	return tokenStr, err
}

func ParseTokenStr(tokenStr string, authUser *AuthUser) error {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.AuthKey), nil
	}

	token, err := jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return err
	}
	claims := token.Claims.(jwt.MapClaims)
	data := claims["data"].(map[string]interface{})
	authUser.User = data["user"].(string)
	authUser.Password = data["password"].(string)
	return nil

}
