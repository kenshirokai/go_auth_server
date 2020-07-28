package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/kenshirokai/go_app_server/domain"
)

type IClientRepository interface {
	Create(client *domain.Client) error
	FindById(clientID string) (domain.Client, error)
}

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return ClientRepository{
		db: db,
	}
}

func (repository ClientRepository) Create(client *domain.Client) error {
	return repository.db.Create(client).Error
}

func (repository ClientRepository) FindById(clientID string) (domain.Client, error) {
	client := domain.Client{}
	var err error
	err = repository.db.Where("client_id = ?", clientID).First(&client).Error
	return client, err
}
