package server

import (
	"fmt"
	"time"

	"gitlab.com/dvision/jwt-go"
)

const day = 1440

type Claims struct {
	UserId string `json:"userid"`
	Digest string `json:"digest"`
	jwt.StandardClaims
}

type JWT struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

func (s *Server) jwtWithExpTime(userid string, exp int) (*JWT, error) {
	expirationTime := time.Now().Add(time.Duration(exp) * time.Minute)
	claims := &Claims{
		UserId: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.JWTSecretKey)
	if err != nil {
		return nil, fmt.Errorf("jwtWithExpTime:token.SignedString")
	}
	return &JWT{
		Token:   tokenString,
		Expires: expirationTime,
	}, nil
}
