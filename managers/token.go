package managers

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// Jwt token data
type TokenData struct {
	SessionId string
	Username  string
	Role      string
}

const secret = "fidsfskfposkfpsofk"
const expirationMillis = 30 * 60000

// sign jwt token
func signToken(TokenData TokenData) (string, error) {
	// set token data
	claims := &jwt.MapClaims{
		"exp": time.Now().Add(expirationMillis).Unix(),
		"data": map[string]string{
			"sessionId": TokenData.SessionId,
			"username":  TokenData.Username,
			"role":      TokenData.Role,
		},
	}
	// sign token
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	return token.SignedString(secret)
}

// verify token and return decoded TokenData
func verifyToken(tokenString string) (TokenData, error) {
	TokenData := TokenData{}
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
	TokenData.Role = data["role"].(string)
	return TokenData, nil
}
