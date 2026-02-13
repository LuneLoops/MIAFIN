package handlers

import (
	"net/http"

	"github.com/LuneLoops/MIAFIN/internal/auth"
	"github.com/LuneLoops/MIAFIN/internal/database"
	"github.com/LuneLoops/MIAFIN/internal/services"
	"github.com/LuneLoops/MIAFIN/pkg/models"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	logService *services.LogService
}

func NewAuthHandler(logService *services.LogService) *AuthHandler {
	return &AuthHandler{logService: logService}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Usuario struct {
		ID     uint   `json:"id"`
		Nombre string `json:"nombre"`
		Email  string `json:"email"`
		Rol    string `json:"rol"`
	} `json:"usuario"`
}

// Login maneja el inicio de sesión
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var usuario models.Usuario
	result := database.DB.Where("email = ?", req.Email).First(&usuario)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales inválidas"})
		return
	}

	if !auth.CheckPassword(req.Password, usuario.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales inválidas"})
		return
	}

	token, err := auth.GenerateToken(usuario.ID, usuario.Email, usuario.Rol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generando token"})
		return
	}

	response := LoginResponse{
		Token: token,
	}
	response.Usuario.ID = usuario.ID
	response.Usuario.Nombre = usuario.Nombre
	response.Usuario.Email = usuario.Email
	response.Usuario.Rol = usuario.Rol

	c.JSON(http.StatusOK, response)
}
