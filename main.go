package main

import (
	"log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"apiGO/db"
	"apiGO/db/moke"
	"apiGO/db/sqlite"
	"apiGO/service"
	"apiGO/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Config struct {
	ListenPort string
	SecretKey  []byte
	EnvType    string
}

var config Config

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	config.EnvType = viper.GetString("EnvType")
	config.SecretKey = []byte(viper.GetString("SecretKey"))
	config.ListenPort = viper.GetString("ListenPort")
}

func main() {
	r := gin.Default()
	var db *db.Storage
	log.Println("ENV:", config.EnvType)
	if config.EnvType == "dev" {
		log.Println("create Moke DB")
		db = moke.New()
	} else {
		log.Println("create SQLite DB")
		db = sqlite.New("storage.db")
	}

	secureJWT := util.MiddlJWT(config.SecretKey)
	s := service.New(db, config.SecretKey)
	r.GET("/users/:id", s.GetUser)
	r.POST("/users", s.CreateUser)
	r.GET("/users", s.GetAllUser)
	r.DELETE("/users/:id", secureJWT, s.DeleteUser)
	r.POST("/login", s.Login)
	r.Run(":" + config.ListenPort)

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	// Test handler
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App running")
	})

	log.Fatal(app.Listen(":5000"))
}