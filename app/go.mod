module github.com/kenshirokai/go_app_server

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/ikawaha/kagome.ipadic v1.1.2 // indirect
	github.com/jinzhu/gorm v1.9.14
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/valyala/fasttemplate v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
)

replace github.com/kenshirokai/go_app_server => ./
