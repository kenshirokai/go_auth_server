package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/kenshirokai/go_app_server/services"
	"github.com/kenshirokai/go_app_server/utils"
)

type AuthController struct {
	service services.IAuthNService
}

func NewAuthController(service services.IAuthNService) AuthController {
	return AuthController{
		service: service,
	}
}

func (controller AuthController) Authentication(c *gin.Context) {
	//認証リクエストからパラメーターを取り出す
	authNParams, err := getAuthenticationParams(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	//認証リクエストが有効かどうか判定
	err = controller.service.IsValid(authNParams)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	//認証画面を返す (認証パラメーターをつけてリダイレクト)
	loginPagePath := "/"
	query := url.Values{}
	query.Add("scope", authNParams.Scope)
	query.Add("redirect_uri", authNParams.RedirectURI)
	query.Add("response_type", authNParams.ResponseType)
	query.Add("client_id", authNParams.ClientId)
	query.Add("state", authNParams.State)
	c.Redirect(http.StatusTemporaryRedirect, loginPagePath+"?"+query.Encode())
}

func (controller AuthController) Login(c *gin.Context) {}

func (controller AuthController) Authorization(c *gin.Context) {}

func getAuthenticationParams(c *gin.Context) (utils.AuthenticationRequestDto, error) {
	var params utils.AuthenticationRequestDto
	scope := c.Query("scope")
	if scope == "" {
		return params, errors.New(fmt.Sprintf("scope is required"))
	}
	responseType := c.Query("response_type")
	if responseType == "" {
		return params, errors.New(fmt.Sprintf("response_type is required"))
	}
	clientID := c.Query("client_id")
	if clientID == "" {
		return params, errors.New(fmt.Sprintf("client_id is required"))
	}
	redirectURL := c.Query("redirect_uri")
	if redirectURL == "" {
		return params, errors.New(fmt.Sprintf("redirect_uri is required"))
	}
	state := c.Query("state")
	if state == "" {
		return params, errors.New(fmt.Sprintf("state is required"))
	}
	params.Scope = scope
	params.ResponseType = responseType
	params.ClientId = clientID
	params.RedirectURI = redirectURL
	params.State = state
	return params, nil
}
