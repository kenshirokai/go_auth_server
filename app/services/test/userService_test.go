package services_test

import (
	"testing"

	"github.com/kenshirokai/go_app_server/repositories"
	"github.com/kenshirokai/go_app_server/services"
	"github.com/kenshirokai/go_app_server/utils"
	"github.com/kenshirokai/go_app_server/domain"
)

func TestCreate(t *testing.T) {
	//setup
	service := getUserService()
	cases := []struct {
		Input utils.UserCreateRequestDto
		Want  error
	}{
		{
			Input: utils.UserCreateRequestDto{
				Name:     "kai",
				Password: "secret",
				Email:    "test@gmail",
			},
			Want: nil,
		},
	}
	//test body
	for _, data := range cases {
		out := service.Create(data.Input)
		if out != data.Want {
			t.Errorf("Want %v but got %s", data.Want, out)
		}

	}

}

func TestFindByEmail(t *testing.T) {
	//setup
	service := getUserService()
	testEmail := "findbyemail@gmail"
	badEmail := "badEmail"
	user := domain.User{
		Name: "test Find By Email",
		Email: testEmail,
		Password: "test FindByEmail",
	}
	if err := testdb.Create(&user).Error; err != nil {
		panic(err)
	}
	testCases := []struct {
		Input string
		Want uint
	} {
		{Input: testEmail, Want: user.ID},
		{Input: badEmail, Want: 0},
	}
	//test body
	for _, testData := range testCases {
		resultUser, _ := service.FindByEmail(testData.Input)
		if resultUser.ID != testData.Want {
			t.Errorf("Want %v but got %v", testData.Want, resultUser.ID)
		}
	}
}


/*
  Helpers
*/
func getUserService() services.UserService {
	return services.NewUserService(
		repositories.NewUserRepository(testdb))
}
