package services

import (
	"strings"

	"github.com/kenshirokai/go_app_server/utils"
)

type IAuthNService interface {
	//認証リクエストが有効かどうか
	IsValid(dto utils.AuthenticationRequestDto) bool
}

type AuthNService struct{}

func NewAuthNService() AuthNService {
	return AuthNService{}
}

func (service AuthNService) IsValid(dto utils.AuthenticationRequestDto) bool {
	//scopeにopenIdが含まれているか確認
	if !strings.Contains(dto.Scope, "opentid") {
		return false
	}
	//responseTypeはcodeか
	if dto.ResponseType == "code" {
		return false
	}
	//clientIDが登録されているか確認しクライアントを特定
	//特定したclientとclientが登録しているredirectUrlが一致しているかを確認
	return true
}
