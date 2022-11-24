package handler

import (
	"errors"
	"fmt"
	"go-bitly/pkg/db"
	"go-bitly/pkg/models"
	"go-bitly/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllBitlys(c *gin.Context) {
	var bitlys []models.Bitly

	tx := db.DB.Find(&bitlys)
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": tx.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bitlys,
	})
}

func GetBitlyById(c *gin.Context) {
	id := c.Param("id")
	var bitly models.Bitly
	result := db.DB.First(&bitly.Bitly, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Could not find bitly with id %v", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bitly": result,
	})
}

func CreateBitly(c *gin.Context) {
	var payload struct {
		Bitly    string
		Redirect string
		Random   bool
	}

	fmt.Printf("new bitly %+v\n", &payload)

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Invalid JSON body",
		})
		return
	}

	newBitly := models.Bitly{Bitly: payload.Bitly, Redirect: payload.Redirect, Random: payload.Random, Clicked: 0}
	if newBitly.Random {
		newBitly.Bitly = utils.RandomURL(8)
	}

	result := db.DB.Create(&newBitly)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error creating bitly %+v\n", result.Error),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bitly": newBitly,
	})
}
