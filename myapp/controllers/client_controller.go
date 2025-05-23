package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapp/models"
)

func GetClients(c *gin.Context, db *gorm.DB) {
	var clients []models.Client
	db.Find(&clients)
	c.JSON(http.StatusOK, clients)
}

func CreateClient(c *gin.Context, db *gorm.DB) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&client)
	c.JSON(http.StatusOK, client)
}