package repositories

import (
	"testing"

	"github.com/kenshirokai/go_app_server/domain"
)

func TestClientCreate(t *testing.T) {
	//setup
	clientRepostiry := NewClientRepository(testdb)
	testCases := []struct {
		Input domain.Client
		Want  error
	}{
		{
			Input: domain.Client{
				Name:        "test client",
				ClientID:    "test clientID",
				RedirectURI: "https://xxx.yyyy.com",
			},
			Want: nil,
		},
	}
	//test body
	for _, testData := range testCases {
		out := clientRepostiry.Create(&testData.Input)
		if out != testData.Want {
			t.Errorf("Want %v but got %v", testData.Want, out)
		}
	}
}

func TestClientFindById(t *testing.T) {
	clientID := "testClientID"
	clientRepository := NewClientRepository(testdb)
	testClient := domain.Client{
		ClientID:    clientID,
		RedirectURI: "https://findbyid.xxx.com",
		Name:        "test client",
	}
	err := testdb.Create(&testClient).Error
	if err != nil {
		panic(err)
	}
	testCases := []struct {
		Input string
		Want  string
	}{
		{Input: clientID, Want: clientID},
		{Input: "bad id", Want: ""},
	}
	//test body
	for _, testData := range testCases {
		client, _ := clientRepository.FindById(testData.Input)
		if client.ClientID != testData.Want {
			t.Errorf("Want %v but got %v", testData.Want, client.ClientID)
		}
	}
}
