package managers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mattiabonardi/http-backend-go/types"
)

// print to http response internal server error
func ThrowInternalServerError(c *gin.Context, err error) {
	// create response
	ApplicationErrorResponse := types.ApplicationErrorResponse{}
	ApplicationErrorResponse.Status = http.StatusInternalServerError
	ApplicationErrorResponse.Message = err.Error()
	c.AbortWithStatusJSON(http.StatusInternalServerError, ApplicationErrorResponse)
}

// print to http response bad request error
func ThrowBadRequest(c *gin.Context, err error) {
	// create response
	ApplicationErrorResponse := types.ApplicationErrorResponse{}
	ApplicationErrorResponse.Status = http.StatusBadRequest
	ApplicationErrorResponse.Message = err.Error()
	c.AbortWithStatusJSON(http.StatusBadRequest, ApplicationErrorResponse)
}

// print to http response unauthorized error
func ThrowUnauthorize(c *gin.Context, err error) {
	// create response
	ApplicationErrorResponse := types.ApplicationErrorResponse{}
	ApplicationErrorResponse.Status = http.StatusUnauthorized
	ApplicationErrorResponse.Message = err.Error()
	c.AbortWithStatusJSON(http.StatusUnauthorized, ApplicationErrorResponse)
}
