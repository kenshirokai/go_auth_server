package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kenshirokai/go_app_server/domain"
)

var db *gorm.DB

func Init() {
	connect()
	migrate()
}

func GetDbInstance() *gorm.DB {
	return db
}

func connect() {
	var err error
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PASS"), os.Getenv("DB_MODE")))
	if err != nil {
		panic(err)
	}
}
func migrate() {
	db.AutoMigrate(domain.User{})
}
