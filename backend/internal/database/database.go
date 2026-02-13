package database

import (
	"fmt"
	"log"

	"github.com/LuneLoops/MIAFIN/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect establece la conexión con la base de datos
func Connect(host, user, password, dbname, port string) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/La_Paz",
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")
	return nil
}

// Migrate ejecuta las migraciones automáticas
func Migrate() error {
	err := DB.AutoMigrate(
		&models.Usuario{},
		&models.Cliente{},
		&models.Grupo{},
		&models.Prestamo{},
		&models.Pago{},
		&models.LogTransaccion{},
	)

	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migrated successfully")
	return nil
}
