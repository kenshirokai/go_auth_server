package repositories

import (
	"testing"

	"github.com/kenshirokai/go_app_server/utils"
)

//test data
var (
	code        = "testcode"
	state       = "teststate"
	redirectURI = "http://testhost"
)

func TestFindByCode(t *testing.T) {
	setupFindByCode()
	repo := NewAuthRepository(pool)
	testcases := []struct {
		Input string
		Want  utils.AuthFlowInfo
	}{
		{Input: code, Want: utils.AuthFlowInfo{
			State:       state,
			RedirectURI: redirectURI,
		}},
		{Input: "failcode", Want: utils.AuthFlowInfo{}},
	}
	//test body
	for _, test := range testcases {
		result, _ := repo.FindByCode(test.Input)
		if result.RedirectURI != test.Want.RedirectURI && result.State != test.Want.RedirectURI {
			t.Errorf("want state=%s redirect_uri=%s but got state=%s redirect_uri=%s",
				test.Want.State, test.Want.RedirectURI, result.State, result.RedirectURI)
		}
	}
}

func TestSetCodeValues(t *testing.T) {
	repo := NewAuthRepository(pool)
	testcases := []struct {
		Input utils.AuthFlowInfo
		Want  error
	}{
		{Input: utils.AuthFlowInfo{
			Code:        "setTestCode",
			State:       "setTestState",
			RedirectURI: "http://setTestHost",
		}, Want: nil},
	}
	//test body
	for _, test := range testcases {
		err := repo.SetCodeValues(test.Input.Code, test.Input)
		if err != test.Want {
			t.Errorf("want %v but bot %v", test.Want, err)
		}
	}
}

/*
	Helpers
*/
func setupFindByCode() {
	conn := pool.Get()
	conn.Do("HSET", code, "state", state)
	conn.Do("HSET", code, "redirect_uri", redirectURI)
}
