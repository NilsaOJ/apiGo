package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"apiGO/models"
)

// go to Service folder.
func (s *Service) GetRecipe(c *gin.Context) {
	id := c.Param("id")
	u, err := s.db.Recipe.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"id": id,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}

func (s *Service) GetAllRecipe(c *gin.Context) {
	re, err := s.db.Recipe.GetAll()
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"recipes": re,
	})
}

func (s *Service) CreateRecipe(c *gin.Context) {
	var r models.Recipe
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	if len(r.Name) == 0 || len(r.Details) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "need name and details",
		})
		return
	}

	_, err = s.db.Recipe.Create(&r)
	if err != nil {
		log.Println("service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error internal",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"recipes": r,
	})
}

func (s *Service) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error id": id,
		})
		return
	}
	err := s.db.Recipe.DeleteByID(id)
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

func (s *Service) Name(c *gin.Context) {
	var r models.Recipe
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	re, err := s.db.Recipe.GetByName(r.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"name": r.Name,
		})
		return
	}

	if re.Name != r.Name {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "not authorized",
		})
		return
	}
}


