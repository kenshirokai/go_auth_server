package repositories_test

import (
	"fmt"
	"os"
	"testing"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kenshirokai/go_app_server/domain"
)

var testdb *gorm.DB

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func setUp() {
	var err error
	godotenv.Load("../../_test/test.env")
	testdb, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_MODE")))
	if err != nil {
		panic(err)
	}
	if err := testdb.DB().Ping(); err != nil {
		panic(err)
	}
	testdb.AutoMigrate(&domain.User{})
}

func tearDown() {
	testdb.DropTableIfExists(&domain.User{})
}
