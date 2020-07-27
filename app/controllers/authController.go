package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kenshirokai/go_app_server/utils"
)

type AuthController struct{}

func NewAuthController() AuthController {
	return AuthController{}
}

func (controller AuthController) Authentication(c *gin.Context) {

	authNParams, err := getAuthenticationParams(c)
	fmt.Println(authNParams)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"result": authNParams,
	})

}

func getAuthenticationParams(c *gin.Context) (utils.AuthenticationRequestDto, error) {
	var params utils.AuthenticationRequestDto
	var scope string
	if scope = c.Query("scope"); scope == "" {
		return params, errors.New(fmt.Sprintf("scope is required"))
	}
	var responseType string
	if responseType = c.Query("response_type"); responseType == "" {
		return params, errors.New(fmt.Sprintf("response_type is required"))
	}
	var clientID string
	if clientID = c.Query("client_id"); clientID == "" {
		return params, errors.New(fmt.Sprintf("client_id is required"))
	}
	var redirectURL string
	if redirectURL = c.Query("redirect_url"); redirectURL == "" {
		return params, errors.New(fmt.Sprintf("redirect_url is required"))
	}
	var state string
	if state = c.Query("state"); state == "" {
		return params, errors.New(fmt.Sprintf("state is required"))
	}
	params.Scope = scope
	params.ResponseType = responseType
	params.ClientId = clientID
	params.RedirectURI = redirectURL
	params.State = state
	return params, nil
}
