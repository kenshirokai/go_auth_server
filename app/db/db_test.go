package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(t *testing.M) {
	setUp()
	i := t.Run()
	os.Exit(i)
}

func setUp() {
	err := godotenv.Load(fmt.Sprintf("../%s.env", "develop"))
	if err != nil {
		panic(err)
	}
	Init()
}

func TestConnect(t *testing.T) {
	db = GetDbInstance()
	if err := db.DB().Ping(); err != nil {
		t.Error("not connection")
	}
	db.Close()
}
