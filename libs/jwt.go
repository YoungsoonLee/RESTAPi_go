package libs

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

type EasyToken struct {
	Displayname string
	Uid         int64
	Expires     int64
}

var (
	verifyKey string
	//ErrAbsent  = "token absent"
	//ErrInvalid = "token invalid"
	//ErrExpired = "token expired"
	//ErrOther = "other error"
)

func init() {
	verifyKey = beego.AppConfig.String("jwt::token")
}

func (e EasyToken) GetToken() (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: e.Expires, //time.Unix(c.ExpiresAt, 0)
		Issuer:    e.Displayname,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(verifyKey))
	if err != nil {
		//log.Println(err)
		beego.Error("Get token error: ", err)
	}
	return tokenString, err
}

func (e EasyToken) ValidateToken(tokenString string) (bool, error) {
	if tokenString == "" {
		//return false, errors.New(ErrAbsent)
		return false, errors.New(ErrTokenAbsent.Message)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(verifyKey), nil
	})
	if token == nil {
		//return false, errors.New(ErrInvalid)
		return false, errors.New(ErrTokenInvalid.Message)
	}
	if token.Valid {

		return true, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, errors.New(ErrTokenInvalid.Message)
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, errors.New(ErrTokenInvalid.Message)
		} else {
			return false, errors.New(ErrTokenOther.Message)
		}
	} else {
		return false, errors.New(ErrTokenOther.Message)
	}
}
