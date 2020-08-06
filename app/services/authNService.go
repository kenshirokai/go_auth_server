package services

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kenshirokai/go_app_server/domain"
	"github.com/kenshirokai/go_app_server/repositories"
	"github.com/kenshirokai/go_app_server/utils"
	"golang.org/x/crypto/bcrypt"
)

type IAuthNService interface {
	//認証リクエストが有効かどうか
	IsValid(dto utils.AuthenticationRequestDto) error
	Login(dto utils.LoginRequestDto) (token string, err error)
}

type AuthNService struct {
	clientRepository repositories.IClientRepository
	userRepository   repositories.IUserRepository
}

func NewAuthNService(
	clientRepository repositories.IClientRepository,
	userRepository repositories.IUserRepository) AuthNService {
	return AuthNService{
		clientRepository: clientRepository,
		userRepository:   userRepository,
	}
}

//Errors
var ScopeErr = errors.New("scope should be include openid")
var ResponseTypeErr = errors.New("response_type should be code")
var InvalidClientErr = errors.New("invalid client")
var InvalidRedirectUrlErr = errors.New("invalid redirectURL")

//有効な認証リクエストか
func (service AuthNService) IsValid(dto utils.AuthenticationRequestDto) error {
	//scopeにopenIdが含まれているか確認
	if !strings.Contains(dto.Scope, "openid") {
		return ScopeErr
	}
	//responseTypeはcodeか
	if dto.ResponseType != "code" {
		return ResponseTypeErr
	}
	//clientIDが登録されているか確認しクライアントを特定
	client, err := service.clientRepository.FindById(dto.ClientId)
	if err != nil {
		return InvalidClientErr
	}
	//特定したclientとclientが登録しているredirectUrlが一致しているかを確認
	if client.RedirectURI != dto.RedirectURI {
		return InvalidRedirectUrlErr
	}
	return nil
}

//loginリクエストをもとに認証していいユーザーか判定する
func (service AuthNService) Login(dto utils.LoginRequestDto) (token string, err error) {
	user, err := service.userRepository.FindByEmail(dto.Email)
	if err != nil {
		err = errors.New("登録されていないemailです")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
	if err != nil {
		err = errors.New("emailもしくはpasswordが間違っています")
		return
	}
	return
}

func createToken(user domain.User) (tokenStr string, err error) {
	//head
	token := jwt.New(jwt.SigningMethodHS256)
	//payload
	cliems := token.Claims.(jwt.MapClaims)
	cliems["email"] = user.Email
	cliems["name"] = user.Name
	cliems["iat"] = time.Now()
	cliems["exp"] = time.Now().Add(time.Hour * 24).Unix()
	//sign
	tokenStr, err = token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return
}
