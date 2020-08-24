package db

import (
	"fmt"
	"os"

	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kenshirokai/go_app_server/domain"
)

var db *gorm.DB
var pool *redis.Pool

func Init() {
	connect()
	migrate()
}

func GetDbInstance() *gorm.DB {
	return db
}
func GetPool() *redis.Pool {
	return pool
}
func connect() {
	var err error
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PASS"), os.Getenv("DB_MODE")))
	if err != nil {
		panic(err)
	}
	//redis
	pool = redis.NewPool(func() (redis.Conn, error) { return redis.Dial("tcp", os.Getenv("REDIS_ADDR")) }, 10)
}
func migrate() {
	db.AutoMigrate(&domain.User{}, &domain.Client{})
}
