package services

import (
	"fmt"
	"testing"

	"github.com/kenshirokai/go_app_server/domain"
	"github.com/kenshirokai/go_app_server/repositories"
	"github.com/kenshirokai/go_app_server/utils"
)

func TestIsValid(t *testing.T) {
	//setup
	authNService := getAuthNService()
	clientID := "testClientID"
	redirectURL := "http://isvalid.xxx.com"
	client := domain.Client{
		ClientID:    clientID,
		RedirectURI: redirectURL,
		Name:        "gen app",
	}
	if err := testdb.Create(&client).Error; err != nil {
		panic(err)
	}
	testCases := []struct {
		Input utils.AuthenticationRequestDto
		Want  error
	}{
		//case: success
		{
			Input: utils.AuthenticationRequestDto{
				Scope:        "openid",
				ResponseType: "code",
				RedirectURI:  redirectURL,
				ClientId:     clientID,
				State:        "",
			},
			Want: nil,
		},
		//case: failure ※todo時間があるときにしっかり失敗ケース書く
		{
			Input: utils.AuthenticationRequestDto{
				Scope:        "open_id",
				ResponseType: "NoCode",
				RedirectURI:  "https://google.com",
				ClientId:     "gensan",
				State:        "",
			},
			Want: ScopeErr,
		},
	}
	//test body
	for _, testData := range testCases {
		out := authNService.IsValid(testData.Input)
		if out != testData.Want {
			t.Errorf("Want %v but got %v", testData.Want, out)
		}
	}
}

func TestGenerateCode(t *testing.T) {
	testcases := []struct {
		Input int
		Want  int
	}{
		{Input: 20, Want: 20},
		{Input: 100, Want: 100},
		{Input: 150, Want: 150},
	}
	//test body
	for _, test := range testcases {
		code := generateCode(test.Input)
		fmt.Printf("generated code : %s\r\n", code)
		if len(code) != test.Want {
			t.Errorf("want %v but got %v", test.Want, code)
		}
	}
}

/*
  Helpers
*/
func getAuthNService() AuthNService {
	return NewAuthNService(
		repositories.NewClientRepository(testdb),
		repositories.NewUserRepository(testdb),
		repositories.NewAuthRepository(pool))
}
