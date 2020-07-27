package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kenshirokai/go_app_server/db"
)

func Configure() {
	env("develop")
	db.Init()
}

func env(mode string) {
	err := godotenv.Load(fmt.Sprintf("./%s.env", mode))
	if err != nil {
		panic(err)
	}
}
