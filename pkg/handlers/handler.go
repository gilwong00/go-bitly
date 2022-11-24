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
	result := db.DB.First(&bitly, id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Could not find bitly with id %v", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"bitly": bitly,
	})
}

func CreateBitly(c *gin.Context) {
	var payload struct {
		Bitly    string
		Redirect string
		Random   bool
	}

	// fmt.Printf("new bitly %+v\n", &payload)

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func UpdateBitly(c *gin.Context) {
	type Payload struct {
		Bitly    string `form:"bitly" json:"bitly"`
		Redirect string `form:"redirect" json:"redirect"`
	}
	id := c.Param("id")
	var json Payload

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated := map[string]interface{}{}

	if len(json.Bitly) != 0 {
		updated["bitly"] = json.Bitly
	}

	if len(json.Redirect) != 0 {
		updated["redirect"] = json.Redirect
	}

	var bitly models.Bitly
	db.DB.First(&bitly, id)
	db.DB.Model(&bitly).Updates(updated)

	c.JSON(http.StatusOK, gin.H{
		"bitly": bitly,
	})
}
