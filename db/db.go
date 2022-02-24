package db

import "apiGO/models"

type Storage struct {
	User StorageUser
	Ingredient StorageIngredient
	Recipe StorageRecipe
}

type StorageUser interface {
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]models.User, error)
	DeleteByID(id string) error
	Create(u *models.User) (*models.User, error)
	Update(id string, data map[string]interface{}) (*models.User, error)
}

type StorageRecipe interface {
	GetByID(id string) (*models.Recipe, error)
	GetByName(name string) (*models.Recipe, error)
	GetAll() ([]models.Recipe, error)
	DeleteByID(id string) error
	Create(r *models.Recipe) (*models.Recipe, error)
	Update(id string, data map[string]interface{}) (*models.Recipe, error)
}

type StorageIngredient interface {
	GetByID(id string) (*models.Ingredient, error)
	GetByName(email string) (*models.Ingredient, error)
	GetAll() ([]models.Ingredient, error)
	DeleteByID(id string) error
	Create(i *models.Ingredient) (*models.Ingredient, error)
	Update(id string, data map[string]interface{}) (*models.Ingredient, error)
}