package controller

import (
	"marketplace-soa/model"
	"marketplace-soa/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	service service.ClientService
}

func NewClientController(s service.ClientService) *ClientController {
	return &ClientController{service: s}
}

// @Summary Cria um novo cliente
// @Description Cria um cliente com os dados enviados no corpo da requisição
// @Tags clients
// @Accept json
// @Produce json
// @Param client body model.ClientRequest true "Dados do cliente"
// @Success 201 {object} model.Client
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clients/ [post]
func (cc *ClientController) Create(c *gin.Context) {
	var req model.ClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := req.Client()
	id, err := cc.service.Create(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	client.ID = id
	c.JSON(http.StatusCreated, client)
}

// @Summary Atualiza um cliente existente
// @Description Atualiza os dados de um cliente com base no ID
// @Tags clients
// @Accept json
// @Produce json
// @Param id path int true "ID do cliente"
// @Param client body model.ClientRequest true "Dados do cliente"
// @Success 200 {object} model.Client
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clients/{id} [put]
func (cc *ClientController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req model.ClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := req.Client()
	client.ID = id

	if err := cc.service.Update(client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, client)
}

// @Summary Deleta um cliente
// @Description Remove um cliente do sistema pelo ID
// @Tags clients
// @Produce json
// @Param id path int true "ID do cliente"
// @Success 204 "No Content"
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /clients/{id} [delete]
func (cc *ClientController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := cc.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// @Summary Busca cliente por ID
// @Description Retorna os dados de um cliente específico
// @Tags clients
// @Produce json
// @Param id path int true "ID do cliente"
// @Success 200 {object} model.Client
// @Failure 400 {object} model.ErrorResponse
// @Failure 404 {object} model.ErrorResponse
// @Router /clients/{id} [get]
func (cc *ClientController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	client, err := cc.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, client)
}

// @Summary Lista todos os clientes
// @Description Retorna todos os clientes cadastrados
// @Tags clients
// @Produce json
// @Success 200 {array} model.Client
// @Success 204 "No Content"
// @Failure 500 {object} model.ErrorResponse
// @Router /clients/ [get]
func (cc *ClientController) GetAll(c *gin.Context) {
	clients, err := cc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(clients) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, clients)
}
