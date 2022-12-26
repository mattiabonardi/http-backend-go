package managers

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mattiabonardi/http-backend-go/types"
)

var jwtKey = []byte("my_secret_key")

func SignAccessToken(TokenData types.TokenData) (string, error) {
	return signToken(TokenData, 30, "accessToken")
}

func SignRefreshToken(TokenData types.TokenData) (string, error) {
	return signToken(TokenData, 120, "refreshToken")
}

// sign jwt token
func signToken(TokenData types.TokenData, expiration time.Duration, tokenType string) (string, error) {
	// set token data
	claims := &jwt.MapClaims{
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": time.Now().Add(expiration * time.Minute).Unix(),
		"data": map[string]string{
			"sessionId": TokenData.SessionId,
			"username":  TokenData.Username,
			"type":      tokenType,
		},
	}
	// sign token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyAccessToken(tokenString string) (types.TokenData, error) {
	return verifyToken(tokenString, "accessToken")
}

func VerifyRefreshToken(tokenString string) (types.TokenData, error) {
	return verifyToken(tokenString, "refreshToken")
}

// verify token and return decoded TokenData
func verifyToken(tokenString string, tokenType string) (types.TokenData, error) {
	TokenData := types.TokenData{}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return TokenData, err
	}
	data := claims["data"].(map[string]interface{})
	TokenData.SessionId = data["sessionId"].(string)
	TokenData.Username = data["username"].(string)
	tType := data["type"].(string)
	if strings.Compare(tType, tokenType) == 0 {
		return TokenData, nil
	} else {
		return TokenData, errors.New("token verification error")
	}
}
