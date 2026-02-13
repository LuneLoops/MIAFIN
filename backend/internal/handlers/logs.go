package handlers

import (
	"net/http"

	"github.com/LuneLoops/MIAFIN/internal/services"
	"github.com/gin-gonic/gin"
)

type LogHandler struct {
	logService *services.LogService
}

func NewLogHandler(logService *services.LogService) *LogHandler {
	return &LogHandler{logService: logService}
}

// GetLogs obtiene todos los logs
func (h *LogHandler) GetLogs(c *gin.Context) {
	logs, err := h.logService.GetLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error obteniendo logs"})
		return
	}

	c.JSON(http.StatusOK, logs)
}
