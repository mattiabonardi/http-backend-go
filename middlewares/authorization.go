package middlewares

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-http-utils/headers"
	"github.com/mattiabonardi/http-backend-go/managers"
)

// get access token and verify user session
func AuthorizationMiddleware(c *gin.Context) {
	// get token from request
	token := c.Request.Header.Get(headers.Authorization)
	if strings.Compare(token, "") == 0 {
		err := errors.New("missing " + headers.Authorization + " header")
		managers.ThrowBadRequest(c, err)
		return
	}
	// remove Bearer constant
	token = strings.ReplaceAll(token, "Bearer ", "")
	// verify token
	TokenData, err := managers.VerifyAccessToken(token)
	if err != nil {
		managers.ThrowUnauthorize(c, err)
		return
	}
	// set token data to context
	c.Set("tokenData", TokenData)
	c.Next()
}
