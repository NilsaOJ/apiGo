package sqlite

import (
	"apiGO/db"
	"apiGO/models"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	Conn *gorm.DB
}

func New(dbName string) *db.Storage {
	conn, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	return &db.Storage{
		User: &SQLite{
			Conn: conn,
		},
	}
}

func (c *SQLite) GetByID(id string) (*models.User, error) {
	var u models.User
	err := c.Conn.First(&u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *SQLite) GetByEmail(email string) (*models.User, error) {
	var u models.User
	err := c.Conn.First(&u, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *SQLite) GetAll() ([]models.User, error) {
	var us []models.User
	err := c.Conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (c *SQLite) DeleteByID(id string) error {
	return c.Conn.Where("id = ?", id).Delete(&models.User{}).Error
}

func (c *SQLite) Create(u *models.User) (*models.User, error) {
	u.Id = uuid.NewString()
	err := c.Conn.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (c *SQLite) Update(id string, data map[string]interface{}) (*models.User, error) {
	u := models.User{Id: id}
	err := c.Conn.Model(&u).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return c.GetByID(id)
}
