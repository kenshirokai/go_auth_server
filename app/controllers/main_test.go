package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kenshirokai/go_app_server/domain"
	"github.com/kenshirokai/go_app_server/repositories"
	"github.com/kenshirokai/go_app_server/services"
)

var testServer *httptest.Server
var testdb *gorm.DB

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	gin.SetMode("debug")
	env()
	dbOpen()
	testServer = httptest.NewServer(getEngine())
}

func tearDown() {
	testServer.Close()
	testdb.DropTableIfExists(&domain.User{})
}

/*
	TestHelpers
*/
func env() {
	err := godotenv.Load("../_test/test.env")
	if err != nil {
		panic(err)
	}
}
func dbOpen() {
	var err error
	testdb, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_MODE")))
	if err != nil {
		panic(err)
	}
	if err = testdb.DB().Ping(); err != nil {
		panic(err)
	}
	testdb.AutoMigrate(&domain.User{})
}
func getEngine() http.Handler {
	engine := gin.New()
	userGroup := engine.Group("/users")
	{
		userservice := services.NewUserService(
			repositories.NewUserRepository(
				testdb))
		usersController := NewUsersController(userservice)
		userGroup.POST("", usersController.Create)
	}
	authGroup := engine.Group("/auth")
	{
		authController := NewAuthController()
		authGroup.GET("/login", authController.Authentication)
	}
	return engine
}
