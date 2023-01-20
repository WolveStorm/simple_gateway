package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const TimeExpiresDuration = time.Hour * 24 * 7

var salt = []byte("lsy520")

type MyClaim struct {
	*jwt.StandardClaims
	AppId string
}

func GenerateToken(AppId string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TimeExpiresDuration).Unix(),
			Issuer:    "gateway",
		},
		AppId: AppId,
	})
	return claims.SignedString(salt)
}

func VerifyToken(token string) (*MyClaim, error) {
	t, err := jwt.ParseWithClaims(token, &MyClaim{}, func(token *jwt.Token) (interface{}, error) {
		return salt, nil
	})
	if err != nil {
		return nil, err
	}
	if myClaim, ok := t.Claims.(*MyClaim); ok {
		return myClaim, nil
	} else {
		return nil, errors.New("解析claim失败")
	}
}
