package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/kenshirokai/go_app_server/domain"
)

type IUserRepository interface {
	Create(user *domain.User) error
	FindByEmail(email string) (domain.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (repo UserRepository) Create(user *domain.User) error {
	err := repo.db.Create(user).Error
	return err
}

func (repo UserRepository) FindByEmail(email string) (domain.User, error) {
	user := domain.User{}
	err := repo.db.Where("email = ?", email).First(&user).Error
	return user, err
}
