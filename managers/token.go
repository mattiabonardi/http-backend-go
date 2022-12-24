package managers

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mattiabonardi/http-backend-go/types"
)

const secret = "fidsfskfposkfpsofk"
const accessTokenexpirationMillis = 30 * 60000
const refreshTokenexpirationMillis = 120 * 60000

func SignAccessToken(TokenData types.TokenData) (string, error) {
	return signToken(TokenData, accessTokenexpirationMillis, "access")
}

func SignRefreshToken(TokenData types.TokenData) (string, error) {
	return signToken(TokenData, refreshTokenexpirationMillis, "refresh")
}

// sign jwt token
func signToken(TokenData types.TokenData, expiration int, tokenType string) (string, error) {
	// set token data
	claims := &jwt.MapClaims{
		"exp": time.Now().Add(time.Duration(expiration)).Unix(),
		"data": map[string]string{
			"sessionId": TokenData.SessionId,
			"username":  TokenData.Username,
			"type":      tokenType,
		},
	}
	// sign token
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(secret)
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
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return TokenData, err
	}
	claims := token.Claims.(jwt.MapClaims)
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
