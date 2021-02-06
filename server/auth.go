package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jekabolt/quotes/signature"
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

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func (s *Server) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	if tokenString == "" {
		return nil, fmt.Errorf("verifyToken:bad auth header")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.JWTSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s *Server) TokenValid(r *http.Request) error {
	token, err := s.VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return fmt.Errorf("TokenValid:bad claims or invalid token")
	}
	return nil
}

func (s *Server) VerifyClaimSignature(mc jwt.MapClaims) (string, error) {
	userid, ok := mc["userid"].(string)
	if !ok {
		return "", fmt.Errorf("VerifyClaimSignature:bad userid")
	}
	digest, ok := mc["digest"].(string)
	if !ok {
		return "", fmt.Errorf("VerifyClaimSignature:bad digest")
	}

	decrypted, err := signature.DecryptMessage(s.AuthSecretKey, digest)
	if err != nil {
		return "", fmt.Errorf("VerifyClaimSignature:DecryptMessage:err[%v]", err.Error())
	}
	if decrypted != userid {
		return "", fmt.Errorf("bad signature comparison")
	}
	return userid, nil

}

func (s *Server) ExtractTokenMetadata(r *http.Request) (*Claims, error) {
	token, err := s.VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("ExtractTokenMetadata:token invalid")
	}

	userid, err := s.VerifyClaimSignature(claims)
	if err != nil {
		return nil, fmt.Errorf("ExtractTokenMetadata:VerifyClaimSignature:err[%v]", err.Error())
	}
	return &Claims{
		UserId: userid,
	}, nil
}
