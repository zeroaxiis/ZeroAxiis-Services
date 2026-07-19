package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtExpiration = 15 * time.Minute


// this will ensure that the jwt is in valid structure 
type JWTClaims struct {
	AdminID  string `json:"admin_id"`
	SessionID string `json:"session_id"`
	jwt.RegisteredClaims
}


// token generation 
func GenerateJWT(adminID, sessionID, secret string) (string, error) {
	claims := JWTClaims{
		AdminID:  adminID,
		SessionID: sessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(jwtExpiration),
			),
			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(secret))
}


//token verification
func VerifyJWT(tokenString, secret string) (*JWTClaims, error) {
	claims := &JWTClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}