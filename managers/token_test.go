package managers

import (
	"strings"
	"testing"
)

func signAndVerifyTokenTest(t *testing.T) {
	TokenData := TokenData{
		SessionId: "xxxx-xxxx-xxxx-xxxx",
		Username:  "admin",
		Role:      "admin",
	}
	tokenString, err := signToken(TokenData)
	if len(tokenString) < 10 || err != nil {
		t.Fatalf("Incorrect token signature")
	}
	TK, err := verifyToken(tokenString)
	if err != nil {
		t.Fatalf("Unable to encode token")
	}
	if strings.Compare(TK.SessionId, TokenData.SessionId) != 0 {
		t.Fatalf("SessionId not decoded")
	}
	if strings.Compare(TK.Username, TokenData.Username) != 0 {
		t.Fatalf("Username not decoded")
	}
	if strings.Compare(TK.Role, TokenData.Role) != 0 {
		t.Fatalf("Role not decoded")
	}
}
