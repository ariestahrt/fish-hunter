package appjwt

import (
	"errors"
	"fish-hunter/util"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = util.GetConfig("JWT_SECRET")

type JWTClaim struct {
	ID 	 	string   `json:"id"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}

func GenerateToken(id string, roles []string) (tokenString string, err error) {
	fmt.Println("Generating token...")
	fmt.Println("ID: ", id)
	fmt.Println("Roles: ", roles)

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		ID:    		id,
		Roles:    roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(jwtKey))

	return
}

func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)
	
	if !ok {
		return errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}

	return nil
}

func GetJWTPayload(signedToken string) *JWTClaim {
	token, _ := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	claims, _ := token.Claims.(*JWTClaim)
	return claims
}