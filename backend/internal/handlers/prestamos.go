package handlers

import (
	"net/http"

	"github.com/LuneLoops/MIAFIN/internal/database"
	"github.com/LuneLoops/MIAFIN/internal/services"
	"github.com/LuneLoops/MIAFIN/pkg/models"
	"github.com/gin-gonic/gin"
)

type PrestamoHandler struct {
	logService *services.LogService
}

func NewPrestamoHandler(logService *services.LogService) *PrestamoHandler {
	return &PrestamoHandler{logService: logService}
}

type CreatePrestamoRequest struct {
	Tipo        string  `json:"tipo" binding:"required,oneof=comunal individual"`
	ClienteID   *uint   `json:"cliente_id"`
	GrupoID     *uint   `json:"grupo_id"`
	Monto       float64 `json:"monto" binding:"required,gt=0"`
	TasaInteres float64 `json:"tasa_interes" binding:"required,gte=0"`
	PlazoMeses  int     `json:"plazo_meses" binding:"required,gt=0"`
}

// GetPrestamos obtiene todos los préstamos
func (h *PrestamoHandler) GetPrestamos(c *gin.Context) {
	var prestamos []models.Prestamo
	database.DB.Preload("Cliente").Preload("Grupo").Find(&prestamos)
	c.JSON(http.StatusOK, prestamos)
}

// CreatePrestamo crea un nuevo préstamo
func (h *PrestamoHandler) CreatePrestamo(c *gin.Context) {
	var req CreatePrestamoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validar que tenga cliente o grupo según el tipo
	if req.Tipo == "individual" && req.ClienteID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "préstamo individual requiere cliente_id"})
		return
	}
	if req.Tipo == "comunal" && req.GrupoID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "préstamo comunal requiere grupo_id"})
		return
	}

	// Calcular saldo inicial con intereses
	interes := req.Monto * (req.TasaInteres / 100)
	saldoTotal := req.Monto + interes

	prestamo := models.Prestamo{
		Tipo:        req.Tipo,
		ClienteID:   req.ClienteID,
		GrupoID:     req.GrupoID,
		Monto:       req.Monto,
		TasaInteres: req.TasaInteres,
		PlazoMeses:  req.PlazoMeses,
		SaldoActual: saldoTotal,
		Estado:      "aprobado",
	}

	result := database.DB.Create(&prestamo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creando préstamo"})
		return
	}

	userID, _ := c.Get("user_id")
	h.logService.CreateLog(userID.(uint), "crear", "prestamo", prestamo.ID, "Préstamo creado")

	// Cargar relaciones para respuesta
	database.DB.Preload("Cliente").Preload("Grupo").First(&prestamo, prestamo.ID)

	c.JSON(http.StatusCreated, prestamo)
}
