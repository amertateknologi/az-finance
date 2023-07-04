package models

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/amerta-teknologi/az-finance/config"
	"github.com/gin-gonic/gin"
)

type Auth struct {
	AccessToken           string `json:"access_token"`
	AccessTokenExpiresAt  string `json:"access_token_expires_at"`
	PlatformId            string `json:"platform_id"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresAt string `json:"refresh_token_expires_at"`
	SessionId             string `json:"session_id"`
	User                  User   `json:"user"`
}

type User struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
}

type AuthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Auth   `json:"data"`
}

func (model *Auth) Login(c *gin.Context) *AuthResponse {
	// Read the request body
	requestBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading request body")
		return &AuthResponse{}
	}

	response := Post(config.Data.DolphinAddress+"/api/v1/auth/local/callback", "application/x-www-form-urlencoded", bytes.NewBuffer(requestBody))
	body := GetBody(response, &AuthResponse{})

	return body.(*AuthResponse)
}
