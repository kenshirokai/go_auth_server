package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenshirokai/go_app_server/services"
	"github.com/kenshirokai/go_app_server/utils"
)

type UsersController struct {
	service services.IUserService
}

func NewUsersController(service services.IUserService) UsersController {
	return UsersController{
		service: service,
	}
}

func (controller UsersController) Create(c *gin.Context) {
	dto := utils.UserCreateRequestDto{}
	var err error
	if err = c.BindJSON(&dto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"error": err,
		})
		return
	}
	err = controller.service.Create(dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"msg": "ok",
	})
}

func (controller UsersController) UserInfo(c *gin.Context) {

}
