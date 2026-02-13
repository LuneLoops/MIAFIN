package handlers

import (
	"net/http"
	"strconv"

	"github.com/LuneLoops/MIAFIN/internal/services"
	"github.com/gin-gonic/gin"
)

type ReporteHandler struct {
	reporteService *services.ReporteService
	logService     *services.LogService
}

func NewReporteHandler(reporteService *services.ReporteService, logService *services.LogService) *ReporteHandler {
	return &ReporteHandler{
		reporteService: reporteService,
		logService:     logService,
	}
}

// GetLiquidez obtiene la liquidez actual
func (h *ReporteHandler) GetLiquidez(c *gin.Context) {
	liquidez, err := h.reporteService.GetLiquidez()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error calculando liquidez"})
		return
	}

	c.JSON(http.StatusOK, liquidez)
}

// GenerarReporteGeneral genera el reporte financiero general en PDF
func (h *ReporteHandler) GenerarReporteGeneral(c *gin.Context) {
	pdf, err := h.reporteService.GenerarReporteGeneral()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generando reporte"})
		return
	}

	userID, _ := c.Get("user_id")
	h.logService.CreateLog(userID.(uint), "generar", "reporte", 0, "Reporte general generado")

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=reporte-general.pdf")
	c.Data(http.StatusOK, "application/pdf", pdf)
}

// GenerarReportePrestamo genera el reporte de un préstamo en PDF
func (h *ReporteHandler) GenerarReportePrestamo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	pdf, err := h.reporteService.GenerarReportePrestamo(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	h.logService.CreateLog(userID.(uint), "generar", "reporte", uint(id), "Reporte de préstamo generado")

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=reporte-prestamo.pdf")
	c.Data(http.StatusOK, "application/pdf", pdf)
}

// GenerarComprobantePago genera el comprobante de pago en PDF
func (h *ReporteHandler) GenerarComprobantePago(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	pdf, err := h.reporteService.GenerarComprobantePago(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	h.logService.CreateLog(userID.(uint), "generar", "reporte", uint(id), "Comprobante de pago generado")

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=comprobante-pago.pdf")
	c.Data(http.StatusOK, "application/pdf", pdf)
}
