package services

import (
	"errors"
	"strings"

	"github.com/kenshirokai/go_app_server/repositories"
	"github.com/kenshirokai/go_app_server/utils"
)

type IAuthNService interface {
	//認証リクエストが有効かどうか
	IsValid(dto utils.AuthenticationRequestDto) error
}

type AuthNService struct {
	clientRepository repositories.IClientRepository
}

func NewAuthNService(clientRepository repositories.IClientRepository) AuthNService {
	return AuthNService{
		clientRepository: clientRepository,
	}
}

var ScopeErr = errors.New("scope should be include openid")
var ResponseTypeErr = errors.New("response_type should be code")
var InvalidClientErr = errors.New("invalid client")
var InvalidRedirectUrlErr = errors.New("invalid redirectURL")

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
