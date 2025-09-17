package controller

import (
	"net/http"
	"strconv"

	"marketplace-soa/model"
	"marketplace-soa/service"

	"github.com/gin-gonic/gin"
)

// ProductController gerencia endpoints de produtos
type ProductController struct {
	service service.ProductService
}

// NewProductController cria um novo ProductController
func NewProductController(s service.ProductService) *ProductController {
	return &ProductController{service: s}
}

// @Summary Cria um novo produto
// @Description Cria um novo produto com os dados fornecidos
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.Product true "Produto"
// @Success 201 {object} model.Product
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /products [post]
func (pc *ProductController) Create(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := pc.service.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	product.ID = id
	c.JSON(http.StatusCreated, product)
}

// @Summary Atualiza um produto existente
// @Description Atualiza os dados de um produto pelo ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "ID do Produto"
// @Param product body model.Product true "Produto atualizado"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /products/{id} [put]
func (pc *ProductController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = id
	if err := pc.service.Update(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Deleta um produto
// @Description Deleta um produto pelo ID
// @Tags products
// @Param id path int true "ID do Produto"
// @Success 204
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /products/{id} [delete]
func (pc *ProductController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := pc.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Busca produto pelo ID
// @Description Retorna um produto pelo seu ID
// @Tags products
// @Param id path int true "ID do Produto"
// @Success 200 {object} model.Product
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /products/{id} [get]
func (pc *ProductController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	product, err := pc.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary Lista todos os produtos
// @Description Retorna todos os produtos cadastrados
// @Tags products
// @Success 200 {array} model.Product
// @Success 204 "No Content"
// @Failure 500 {object} model.ErrorResponse
// @Router /products [get]
func (pc *ProductController) GetAll(c *gin.Context) {
	products, err := pc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, products)
}
