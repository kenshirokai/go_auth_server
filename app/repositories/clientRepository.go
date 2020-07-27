package repositories

import "github.com/kenshirokai/go_app_server/domain"

type IClientRepository interface {
	FindById(clientID string) (domain.Client, error)
}

type ClientRepository struct{}
