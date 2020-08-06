package application

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kenshirokai/go_app_server/config"
	"github.com/kenshirokai/go_app_server/controllers"
	"github.com/kenshirokai/go_app_server/db"
	"github.com/kenshirokai/go_app_server/middleware"
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
	engine.LoadHTMLGlob("static/*.html")
	engine.Static("static", "static")
	engine.Handle("GET", "/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	engine.Use(middleware.CORS)
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
		authController := controllers.NewAuthController(
			services.NewAuthNService(
				repositories.NewClientRepository(
					db.GetDbInstance()),
				repositories.NewUserRepository(
					db.GetDbInstance())))
		authGroup.GET("", authController.Authentication)
	}
}
