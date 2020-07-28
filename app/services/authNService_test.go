package services

import (
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
		RedirectURL: redirectURL,
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

/*
  Helpers
*/
func getAuthNService() AuthNService {
	return NewAuthNService(
		repositories.NewClientRepository(testdb))
}
