package services

import (
	"github.com/kenshirokai/go_app_server/domain"
	"github.com/kenshirokai/go_app_server/repositories"
	"github.com/kenshirokai/go_app_server/utils"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(dto utils.UserCreateRequestDto) error
	FindByEmail(email string) (domain.User, error)
}


type UserService struct {
	repository repositories.IUserRepository
}

func NewUserService(repository repositories.IUserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

func (service UserService) Create(dto utils.UserCreateRequestDto) error {
	var err error
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := domain.User{}
	user.Password = string(hash)
	user.Name = dto.Name
	user.Email = dto.Email

	err = service.repository.Create(&user)

	return err
}

func (service UserService) FindByEmail(email string) (domain.User, error) {
	return service.repository.FindByEmail(email)
}