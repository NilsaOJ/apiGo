package moke

import (
	"apiGO/db"
	"apiGO/models"

	"errors"
	"github.com/google/uuid"
)

var _ db.StorageUser = &Moke{}

type Moke struct {
	listUser map[string]*models.User
}

func New() *db.Storage {
	return &db.Storage{
		User: &Moke{
			listUser: make(map[string]*models.User),
		},
	}
}

func (m *Moke) GetByID(id string) (*models.User, error) {
	u, ok := m.listUser[id]
	if !ok {
		return nil, errors.New("db user: not found")
	}
	return u, nil
}

func (m *Moke) GetByEmail(email string) (*models.User, error) {
	for k := range m.listUser {
		if m.listUser[k].Email == email {
			return m.listUser[k], nil
		}
	}

	return nil, errors.New("db user: not found")
}

func (m *Moke) DeleteByID(id string) error {
	_, ok := m.listUser[id]
	if !ok {
		return errors.New("db user: not found")
	}
	delete(m.listUser, id)
	return nil
}

func (m *Moke) Create(u *models.User) (*models.User, error) {
	u.Id = uuid.New().String()
	m.listUser[u.Id] = u
	return u, nil
}

func (m *Moke) Update(id string, data map[string]interface{}) (*models.User, error) {
	u, ok := m.listUser[id]
	if !ok {
		return nil, errors.New("db user: not found")
	}
	if value, ok := data["first_name"]; ok {
		u.Firstname = value.(string)
	}
	if value, ok := data["last_name"]; ok {
		u.Firstname = value.(string)
	}
	return nil, nil
}

func (m *Moke) GetAll() ([]models.User, error) {
	us := make([]models.User, len(m.listUser))
	var i int
	for k := range m.listUser {
		if m.listUser[k] != nil {
			us[i] = *m.listUser[k]
		}
		i++
	}
	return us, nil
}
