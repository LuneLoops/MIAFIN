package handlers

import (
	"net/http"
	"time"

	"github.com/LuneLoops/MIAFIN/internal/database"
	"github.com/LuneLoops/MIAFIN/internal/services"
	"github.com/LuneLoops/MIAFIN/pkg/models"
	"github.com/gin-gonic/gin"
)

type PagoHandler struct {
	logService *services.LogService
}

func NewPagoHandler(logService *services.LogService) *PagoHandler {
	return &PagoHandler{logService: logService}
}

type CreatePagoRequest struct {
	PrestamoID uint    `json:"prestamo_id" binding:"required"`
	Monto      float64 `json:"monto" binding:"required,gt=0"`
}

// CreatePago registra un nuevo pago
func (h *PagoHandler) CreatePago(c *gin.Context) {
	var req CreatePagoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar que el préstamo existe
	var prestamo models.Prestamo
	result := database.DB.First(&prestamo, req.PrestamoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "préstamo no encontrado"})
		return
	}

	// Verificar que el monto no exceda el saldo
	if req.Monto > prestamo.SaldoActual {
		c.JSON(http.StatusBadRequest, gin.H{"error": "monto excede saldo actual"})
		return
	}

	// Crear el pago
	pago := models.Pago{
		PrestamoID: req.PrestamoID,
		Monto:      req.Monto,
		FechaPago:  time.Now(),
	}

	result = database.DB.Create(&pago)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creando pago"})
		return
	}

	// Actualizar saldo del préstamo
	nuevoSaldo := prestamo.SaldoActual - req.Monto
	database.DB.Model(&prestamo).Update("saldo_actual", nuevoSaldo)

	// Si saldo es 0, marcar como pagado
	if nuevoSaldo == 0 {
		database.DB.Model(&prestamo).Update("estado", "pagado")
	}

	userID, _ := c.Get("user_id")
	h.logService.CreateLog(userID.(uint), "crear", "pago", pago.ID, "Pago registrado")

	c.JSON(http.StatusCreated, pago)
}
