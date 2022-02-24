package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"apiGO/models"
)

// go to Service folder.
func (s *Service) GetIngredient(c *gin.Context) {
	id := c.Param("id")
	i, err := s.db.Ingredient.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ingredient": i,
	})
}

func (s *Service) GetAllIngredient(c *gin.Context) {
	in, err := s.db.Ingredient.GetAll()
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ingredients": in,
	})
}

func (s *Service) CreateIngredient(c *gin.Context) {
	var i models.Ingredient
	err := c.BindJSON(&i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	if len(i.Name) == 0 || len(i.Dlc) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "need name and dlc",
		})
		return
	}

	_, err = s.db.Ingredient.Create(&i)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ingredient": i,
	})
}

func (s *Service) DeleteIngredient(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error id": id,
		})
		return
	}
	err := s.db.Ingredient.DeleteByID(id)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"delete": id,
	})
}

func (s *Service) Names(c *gin.Context) {
	var i models.Ingredient
	err := c.BindJSON(&i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	in, err := s.db.Recipe.GetByName(i.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"name": i.Name,
		})
		return
	}

	if in.Name != i.Name {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "not authorized",
		})
		return
	}
}