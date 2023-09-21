package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ricoaditya-u/hris-master/db"
	"github.com/ricoaditya-u/hris-master/models"
	"gorm.io/gorm"
)

func LevelIndex(c *gin.Context) {
	var level []models.Level
	err := db.DB.Find(&level).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": level,
	})
}

func LevelCreate(c *gin.Context) {
	var body models.Level

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	level := models.Level{Name: body.Name}
	result := db.DB.Create(&level)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	// Return
	c.JSON(http.StatusOK, gin.H{
		"message": "create success",
	})
}

func LevelShow(c *gin.Context) {
	id := c.Param("id")

	var level models.Level
	err := db.DB.First(&level, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": level,
	})
}

func LevelUpdate(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Get data body
	var body models.Level

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the data
	var level models.Level
	err := db.DB.First(&level, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Update
	db.DB.Model(&level).Updates(models.Level{
		Name: body.Name,
	})

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"data": level,
	})
}

func LevelDelete(c *gin.Context) {
	// Get id
	id := c.Param("id")

	// Delete
	var level models.Level
	err := db.DB.Delete(&level, "ID = ?", id).Error

	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "record not found",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "delete success",
	})
}
