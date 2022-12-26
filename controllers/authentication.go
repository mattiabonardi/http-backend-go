package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mattiabonardi/http-backend-go/managers"
	"github.com/mattiabonardi/http-backend-go/types"
)

type AuthenticationController struct{}

func (h AuthenticationController) Login(c *gin.Context) {
	LoginRequestDTO := types.LoginRequestDTO{}
	LoginResponseDTO := types.LoginResponseDTO{}
	// get body
	if err := c.ShouldBindJSON(&LoginRequestDTO); err != nil {
		managers.ThrowBadRequest(c, err)
		return
	}
	// create token data
	TokenData := types.TokenData{}
	TokenData.SessionId = uuid.New().String()
	TokenData.Username = LoginRequestDTO.Username
	// sign access token
	accessToken, err := managers.SignAccessToken(TokenData)
	if err != nil {
		managers.ThrowUnauthorize(c, err)
		return
	}
	// sign refresh token
	refreshToken, err := managers.SignRefreshToken(TokenData)
	if err != nil {
		managers.ThrowUnauthorize(c, err)
		return
	}
	// create response
	LoginResponseDTO.AccessToken = "Bearer " + accessToken
	LoginResponseDTO.RefreshToken = "Bearer " + refreshToken
	LoginResponseDTO.Message = "Login successfull"
	c.JSON(http.StatusOK, LoginResponseDTO)
}

func (h AuthenticationController) RefreshToken(c *gin.Context) {
	RefreshTokenRequestDTO := types.RefreshTokenRequestDTO{}
	LoginResponseDTO := types.LoginResponseDTO{}
	// get body
	if err := c.ShouldBindJSON(&RefreshTokenRequestDTO); err != nil {
		managers.ThrowBadRequest(c, err)
		return
	}
	// remove Bearer constant
	RefreshTokenRequestDTO.RefreshToken = strings.ReplaceAll(RefreshTokenRequestDTO.RefreshToken, "Bearer ", "")
	// verify refresh token
	TokenData, err := managers.VerifyRefreshToken(RefreshTokenRequestDTO.RefreshToken)
	if err != nil {
		managers.ThrowUnauthorize(c, err)
		return
	}
	// create new session id
	TokenData.SessionId = uuid.New().String()
	// sign new access token
	accessToken, err := managers.SignAccessToken(TokenData)
	if err != nil {
		managers.ThrowUnauthorize(c, err)
		return
	}
	// create response
	LoginResponseDTO.AccessToken = "Bearer " + accessToken
	LoginResponseDTO.RefreshToken = "Bearer " + RefreshTokenRequestDTO.RefreshToken
	LoginResponseDTO.Message = "Refresh token successfull"
	c.JSON(http.StatusOK, LoginResponseDTO)
}
