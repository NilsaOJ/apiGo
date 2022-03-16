package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"apiGO/cache"
	"apiGO/db"
	"apiGO/db/moke"
	"apiGO/db/mysql"
	"apiGO/db/sqlite"
	"apiGO/service"
	"apiGO/util"
)

type Config struct {
	ListenPort string
	SecretKey  []byte
	EnvType    string
	db         struct {
		DBName string
		User   string
		Pass   string
		Port   string
	}
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
	// connect to DB
	config.db.DBName = viper.GetString("db.DBName")
	config.db.User = viper.GetString("db.User")
	config.db.Pass = viper.GetString("db.Pass")
	config.db.Port = viper.GetString("db.Port")
}

func main() {
	r := gin.Default()
	var DB *db.Storage
	log.Println("ENV:", config.EnvType)
	if config.EnvType == "dev" {
		log.Println("create Moke DB")
		DB = moke.New()
	} else if config.EnvType == "pprod" {
		log.Println("connect to an MySQL")
		DB = mysql.New(config.db.DBName, config.db.User, config.db.Pass, config.db.Port)
	} else {
		log.Println("create SQLite DB")
		DB = sqlite.New("storage.db")
	}

	secureJWT := util.MiddlJWT(config.SecretKey)
	c := cache.New()
	cacheMdw := cache.MiddlCache(c)
	s := service.New(DB, c, config.SecretKey)
	r.GET("/users/:id", cacheMdw, s.GetUser)
	r.POST("/users", s.CreateUser)
	r.GET("/users", cacheMdw, s.GetAllUser)
	r.DELETE("/users/:id", secureJWT, s.DeleteUser)
	r.POST("/login", s.Login)

	r.GET("/recipes/:id", cacheMdw, s.GetRecipe)
	r.POST("/recipes", s.CreateRecipe)
	r.GET("/recipes", cacheMdw, s.GetAllRecipe)
	r.DELETE("/recipes/:id", secureJWT, s.DeleteRecipe)
	r.POST("/name", s.Name)

	r.GET("/ingredients/:id", cacheMdw, s.GetIngredient)
	r.POST("/ingredients", s.CreateIngredient)
	r.GET("/ingredients", cacheMdw, s.GetAllIngredient)
	r.DELETE("/ingredients/:id", secureJWT, s.DeleteIngredient)
	r.POST("/names", s.Names)
	err := r.Run(":" + config.ListenPort)
	if err != nil {
		return 
	}
}