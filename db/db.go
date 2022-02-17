package db

import "apiGO/models"

type Storage struct {
	User StorageUser
}

type StorageUser interface {
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]models.User, error)
	DeleteByID(id string) error
	Create(u *models.User) (*models.User, error)
	Update(id string, data map[string]interface{}) (*models.User, error)
}
