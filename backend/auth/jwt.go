package auth

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWTKEY"))

type JWTClaim struct {
	Email     string `json:"email"`
	AuthLevel uint8  `json:"authLevel"`
	UserID    uint64 `json:"userID"`
	jwt.StandardClaims
}

func GenerateJWT(email string, authLevel uint8, userID uint64) (string, error) {

	// Set the expiration time for the token (1 hour from now)
	expiresAt := time.Now().Add(time.Hour * 1)
	// Using jwt.NewWithClaims to create a new token with specific claims
	claims := &JWTClaim{
		Email:     email,
		AuthLevel: authLevel,
		UserID:    userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(jwtKey)

	return signedToken, err
}

func ValidateToken(signedToken string) (*JWTClaim, error) {
	//parse claims
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return nil, err
	}
	//get original
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, err
}
