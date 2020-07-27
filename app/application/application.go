package application

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kenshirokai/go_app_server/config"
	"github.com/kenshirokai/go_app_server/controllers"
	"github.com/kenshirokai/go_app_server/db"
	"github.com/kenshirokai/go_app_server/repositories"
	"github.com/kenshirokai/go_app_server/services"
)

var engine *gin.Engine

//packageの初期化
func init() {
	engine = gin.Default()
}

func Start() {
	config.Configure()
	setHandler()
	engine.Run(os.Getenv("Addr"))
}

func setHandler() {
	userGroup := engine.Group("/users")
	{
		userService := services.NewUserService(
			repositories.NewUserRepository(
				db.GetDbInstance()))
		usersController := controllers.NewUsersController(userService)
		userGroup.POST("", usersController.Create)
	}
	authGroup := engine.Group("/auth")
	{
		authController := controllers.NewAuthController()
		authGroup.GET("/login", authController.Authentication)
	}
}
