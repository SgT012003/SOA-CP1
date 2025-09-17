// @title Marketplace SOA API
// @version 1.0
// @description API de exemplo com Gin + Swagger
// @host localhost:8080
// @BasePath /
package main

import (
	"marketplace-soa/controller"
	"marketplace-soa/service"
	"net/http"

	_ "marketplace-soa/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name
// @contact.url
// @contact.email
func main() {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	clientController := controller.NewClientController(service.NewClientService())
	productController := controller.NewProductController(service.NewProductService())

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoint raiz
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusSeeOther, "/swagger/index.html")
	})

	// URI e handlers para clients
	clients := r.Group("/clients")
	{
		clients.POST("/", clientController.Create)
		clients.PUT("/:id", clientController.Update)
		clients.DELETE("/:id", clientController.Delete)
		clients.GET("/:id", clientController.GetByID)
		clients.GET("/", clientController.GetAll)
	}

	// URI e handlers para products
	products := r.Group("/products")
	{
		products.POST("/", productController.Create)
		products.PUT("/:id", productController.Update)
		products.DELETE("/:id", productController.Delete)
		products.GET("/:id", productController.GetByID)
		products.GET("/", productController.GetAll)
	}

	// Inicia o servidor
	r.Run("0.0.0.0:8080")
}
