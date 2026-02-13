package handlers

import (
	"net/http"

	"github.com/LuneLoops/MIAFIN/internal/database"
	"github.com/LuneLoops/MIAFIN/internal/services"
	"github.com/LuneLoops/MIAFIN/pkg/models"
	"github.com/gin-gonic/gin"
)

type GrupoHandler struct {
	logService *services.LogService
}

func NewGrupoHandler(logService *services.LogService) *GrupoHandler {
	return &GrupoHandler{logService: logService}
}

type CreateGrupoRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion"`
}

// GetGrupos obtiene todos los grupos
func (h *GrupoHandler) GetGrupos(c *gin.Context) {
	var grupos []models.Grupo
	database.DB.Find(&grupos)
	c.JSON(http.StatusOK, grupos)
}

// CreateGrupo crea un nuevo grupo
func (h *GrupoHandler) CreateGrupo(c *gin.Context) {
	var req CreateGrupoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grupo := models.Grupo{
		Nombre:      req.Nombre,
		Descripcion: req.Descripcion,
	}

	result := database.DB.Create(&grupo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creando grupo"})
		return
	}

	userID, _ := c.Get("user_id")
	h.logService.CreateLog(userID.(uint), "crear", "grupo", grupo.ID, "Grupo creado: "+grupo.Nombre)

	c.JSON(http.StatusCreated, grupo)
}
