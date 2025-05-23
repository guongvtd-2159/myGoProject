package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"myapp/controllers"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		controllers.GetUsers(c, db)
	})
	r.POST("/users", func(c *gin.Context) {
		controllers.CreateUser(c, db)
	})

	r.GET("/clients", func(c *gin.Context) {
		controllers.GetClients(c, db)
	})
	r.POST("/clients", func(c *gin.Context) {
		controllers.CreateClient(c, db)
	})

	r.GET("/products", func(c *gin.Context) {
		controllers.GetProducts(c, db)
	})
	r.POST("/products", func(c *gin.Context) {
		controllers.CreateProduct(c, db)
	})

	return r
}
