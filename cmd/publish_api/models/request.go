package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {//payload,加入自己的信息。自己定义的
	ID uint
	Token string
	jwt.StandardClaims
}
