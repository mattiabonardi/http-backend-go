package managers

import (
	"strings"
	"testing"

	"github.com/mattiabonardi/http-backend-go/types"
)

func TestSignAndVerifyAccessToken(t *testing.T) {
	TokenData := types.TokenData{
		SessionId: "xxxx-xxxx-xxxx-xxxx",
		Username:  "admin",
	}

	tokenString, err := SignAccessToken(TokenData)
	if len(tokenString) < 10 || err != nil {
		t.Fatalf("Incorrect token signature: " + err.Error())
	}
	TK, err := VerifyAccessToken(tokenString)
	if err != nil {
		t.Fatalf("Unable to decode token: " + err.Error())
	}
	if strings.Compare(TK.SessionId, TokenData.SessionId) != 0 {
		t.Fatalf("SessionId not decoded")
	}
	if strings.Compare(TK.Username, TokenData.Username) != 0 {
		t.Fatalf("Username not decoded")
	}
}

func TestSignAndVerifyRefreshToken(t *testing.T) {
	TokenData := types.TokenData{
		SessionId: "xxxx-xxxx-xxxx-xxxx",
		Username:  "admin",
	}
	tokenString, err := SignRefreshToken(TokenData)
	if len(tokenString) < 10 || err != nil {
		t.Fatalf("Incorrect token signature")
	}
	TK, err := VerifyRefreshToken(tokenString)
	if err != nil {
		t.Fatalf("Unable to encode token")
	}
	if strings.Compare(TK.SessionId, TokenData.SessionId) != 0 {
		t.Fatalf("SessionId not decoded")
	}
	if strings.Compare(TK.Username, TokenData.Username) != 0 {
		t.Fatalf("Username not decoded")
	}
}
