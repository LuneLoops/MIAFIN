package handlers

import (
	"net/http"

	"github.com/LuneLoops/MIAFIN/internal/database"
	"github.com/LuneLoops/MIAFIN/internal/services"
	"github.com/LuneLoops/MIAFIN/pkg/models"
	"github.com/gin-gonic/gin"
)

type ClienteHandler struct {
	logService *services.LogService
}

func NewClienteHandler(logService *services.LogService) *ClienteHandler {
	return &ClienteHandler{logService: logService}
}

type CreateClienteRequest struct {
	Nombre   string `json:"nombre" binding:"required"`
	CI       string `json:"ci" binding:"required"`
	Telefono string `json:"telefono"`
	GrupoID  *uint  `json:"grupo_id"`
}

// GetClientes obtiene todos los clientes
func (h *ClienteHandler) GetClientes(c *gin.Context) {
	var clientes []models.Cliente
	database.DB.Preload("Grupo").Find(&clientes)
	c.JSON(http.StatusOK, clientes)
}

// CreateCliente crea un nuevo cliente
func (h *ClienteHandler) CreateCliente(c *gin.Context) {
	var req CreateClienteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cliente := models.Cliente{
		Nombre:   req.Nombre,
		CI:       req.CI,
		Telefono: req.Telefono,
		GrupoID:  req.GrupoID,
	}

	result := database.DB.Create(&cliente)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creando cliente"})
		return
	}

	userID, _ := c.Get("user_id")
	h.logService.CreateLog(userID.(uint), "crear", "cliente", cliente.ID, "Cliente creado: "+cliente.Nombre)

	c.JSON(http.StatusCreated, cliente)
}
