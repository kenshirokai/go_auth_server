package repositories

import (
	"testing"

	"github.com/kenshirokai/go_app_server/domain"
)

func TestCreate(t *testing.T) {
	//setup
	userRepo := NewUserRepository(testdb)
	testCases := []struct {
		Input domain.User
		Want  error
	}{
		{
			Input: domain.User{
				Name:     "repo create test",
				Email:    "repo.create@gmail",
				Password: "secret",
			},
			Want: nil,
		},
	}
	//test body
	for _, data := range testCases {
		out := userRepo.Create(&data.Input)
		if out != data.Want {
			t.Errorf("Want %v but got %v", data.Want, out)
		}
	}

}

func TestFindByEmail(t *testing.T) {
	//setup
	userRepositpry := NewUserRepository(testdb)
	email := "findTest@gmail"
	badEmail := "gad@gmail"
	testUser := domain.User{
		Name:     "test user",
		Email:    email,
		Password: "FindTest",
	}
	err := testdb.Create(&testUser).Error
	if err != nil {
		panic(err)
	}
	//Test body
	testCases := []struct {
		Input string
		Want  uint
	}{
		{Input: email, Want: testUser.ID},
		{Input: badEmail, Want: 0},
	}
	for _, testData := range testCases {
		user, _ := userRepositpry.FindByEmail(testData.Input)
		if user.ID != testData.Want {
			t.Errorf("Want ID=%v but got %v", testData.Want, user.ID)
		}
	}

}
