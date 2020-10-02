package config

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kenshirokai/go_app_server/db"
)

func Configure() {
	mode := flag.String("mode", "develop", "起動するモードの設定")
	flag.Parse()
	fmt.Println("-------------mode-------------")
	fmt.Println("------------" + *mode + "------------")
	env(*mode)
	db.Init()
}

func env(mode string) {
	err := godotenv.Load(fmt.Sprintf("./%s.env", mode))
	if err != nil {
		panic(err)
	}
}
