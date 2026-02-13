package database

import (
	"log"

	"github.com/LuneLoops/MIAFIN/internal/auth"
	"github.com/LuneLoops/MIAFIN/pkg/models"
)

// Seed inserta datos iniciales en la base de datos
func Seed() error {
	// Verificar si ya existe el usuario admin
	var count int64
	DB.Model(&models.Usuario{}).Where("email = ?", "admin@miafin.local").Count(&count)

	if count > 0 {
		log.Println("Seed data already exists, skipping...")
		return nil
	}

	// Crear usuario admin
	hashedPassword, err := auth.HashPassword("admin123")
	if err != nil {
		return err
	}

	admin := models.Usuario{
		Nombre:       "Administrador",
		Email:        "admin@miafin.local",
		PasswordHash: hashedPassword,
		Rol:          "admin",
	}

	result := DB.Create(&admin)
	if result.Error != nil {
		return result.Error
	}

	log.Println("Seed data created successfully")
	return nil
}
