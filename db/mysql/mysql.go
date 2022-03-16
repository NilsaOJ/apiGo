package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"apiGO/db"
	"apiGO/db/sqlite"
	"apiGO/models"
)

type MySQL = sqlite.SQLite

func New(dbName, user, pass, port string) *db.Storage {
	// dsn := "miamideas-user:miamideas-pwd@tcp(localhost:3306)/miamideas-db?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%v:%v@tcp(localhost:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, port, dbName)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = conn.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}

	return &db.Storage{
		User: &MySQL{
			Conn: conn,
		},
	}
}
