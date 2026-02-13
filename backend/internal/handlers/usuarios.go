package handlers

import (
	"net/http"

	"github.com/LuneLoops/MIAFIN/internal/auth"
	"github.com/LuneLoops/MIAFIN/internal/database"
	"github.com/LuneLoops/MIAFIN/internal/services"
	"github.com/LuneLoops/MIAFIN/pkg/models"
	"github.com/gin-gonic/gin"
)

type UsuarioHandler struct {
	logService *services.LogService
}

func NewUsuarioHandler(logService *services.LogService) *UsuarioHandler {
	return &UsuarioHandler{logService: logService}
}

type CreateUsuarioRequest struct {
	Nombre   string `json:"nombre" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Rol      string `json:"rol" binding:"required,oneof=admin asesor"`
}

// GetUsuarios obtiene todos los usuarios
func (h *UsuarioHandler) GetUsuarios(c *gin.Context) {
	var usuarios []models.Usuario
	database.DB.Find(&usuarios)
	c.JSON(http.StatusOK, usuarios)
}

// CreateUsuario crea un nuevo usuario
func (h *UsuarioHandler) CreateUsuario(c *gin.Context) {
	var req CreateUsuarioRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error procesando contraseña"})
		return
	}

	usuario := models.Usuario{
		Nombre:       req.Nombre,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		Rol:          req.Rol,
	}

	result := database.DB.Create(&usuario)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creando usuario"})
		return
	}

	userID, _ := c.Get("user_id")
	h.logService.CreateLog(userID.(uint), "crear", "usuario", usuario.ID, "Usuario creado: "+usuario.Email)

	c.JSON(http.StatusCreated, usuario)
}
