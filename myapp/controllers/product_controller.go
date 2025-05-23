package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapp/models"
)

func GetProducts(c *gin.Context, db *gorm.DB) {
	var products []models.Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context, db *gorm.DB) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&product)
	c.JSON(http.StatusOK, product)
}