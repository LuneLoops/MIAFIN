package services

import (
	"github.com/LuneLoops/MIAFIN/internal/database"
	"github.com/LuneLoops/MIAFIN/pkg/models"
)

// LogService maneja las operaciones de logging
type LogService struct{}

// NewLogService crea una nueva instancia de LogService
func NewLogService() *LogService {
	return &LogService{}
}

// CreateLog registra una transacción en el log
func (s *LogService) CreateLog(usuarioID uint, accion, entidad string, entidadID uint, descripcion string) error {
	log := models.LogTransaccion{
		UsuarioID:   usuarioID,
		Accion:      accion,
		Entidad:     entidad,
		EntidadID:   entidadID,
		Descripcion: descripcion,
	}

	result := database.DB.Create(&log)
	return result.Error
}

// GetLogs obtiene todos los logs
func (s *LogService) GetLogs() ([]models.LogTransaccion, error) {
	var logs []models.LogTransaccion
	result := database.DB.Preload("Usuario").Order("created_at DESC").Find(&logs)
	return logs, result.Error
}
