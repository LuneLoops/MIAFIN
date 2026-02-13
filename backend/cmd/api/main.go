package main

import (
	"log"
	"os"

	"github.com/LuneLoops/MIAFIN/internal/database"
	"github.com/LuneLoops/MIAFIN/internal/handlers"
	"github.com/LuneLoops/MIAFIN/internal/middleware"
	"github.com/LuneLoops/MIAFIN/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Configuración de base de datos desde variables de entorno
	dbHost := getEnv("DB_HOST", "localhost")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "miafin")
	dbPort := getEnv("DB_PORT", "5432")

	// Conectar a la base de datos
	err := database.Connect(dbHost, dbUser, dbPassword, dbName, dbPort)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ejecutar migraciones
	err = database.Migrate()
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Ejecutar seed
	err = database.Seed()
	if err != nil {
		log.Printf("Warning: Failed to seed database: %v", err)
	}

	// Inicializar servicios
	logService := services.NewLogService()
	reporteService := services.NewReporteService()

	// Inicializar handlers
	authHandler := handlers.NewAuthHandler(logService)
	usuarioHandler := handlers.NewUsuarioHandler(logService)
	clienteHandler := handlers.NewClienteHandler(logService)
	grupoHandler := handlers.NewGrupoHandler(logService)
	prestamoHandler := handlers.NewPrestamoHandler(logService)
	pagoHandler := handlers.NewPagoHandler(logService)
	logHandler := handlers.NewLogHandler(logService)
	reporteHandler := handlers.NewReporteHandler(reporteService, logService)

	// Configurar Gin
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rutas públicas
	router.POST("/login", authHandler.Login)

	// Rutas protegidas
	api := router.Group("/")
	api.Use(middleware.AuthMiddleware())
	{
		// Usuarios (solo admin)
		usuarios := api.Group("/usuarios")
		usuarios.Use(middleware.RoleMiddleware("admin"))
		{
			usuarios.GET("", usuarioHandler.GetUsuarios)
			usuarios.POST("", usuarioHandler.CreateUsuario)
		}

		// Clientes
		api.GET("/clientes", clienteHandler.GetClientes)
		api.POST("/clientes", clienteHandler.CreateCliente)

		// Grupos (solo admin)
		grupos := api.Group("/grupos")
		{
			grupos.GET("", grupoHandler.GetGrupos)
			grupos.POST("", middleware.RoleMiddleware("admin"), grupoHandler.CreateGrupo)
		}

		// Préstamos
		api.GET("/prestamos", prestamoHandler.GetPrestamos)
		api.POST("/prestamos", prestamoHandler.CreatePrestamo)

		// Pagos
		api.POST("/pagos", pagoHandler.CreatePago)

		// Reportes
		api.GET("/reportes/liquidez", reporteHandler.GetLiquidez)
		api.GET("/reportes/general/pdf", reporteHandler.GenerarReporteGeneral)
		api.GET("/reportes/prestamo/:id/pdf", reporteHandler.GenerarReportePrestamo)
		api.GET("/reportes/pago/:id/pdf", reporteHandler.GenerarComprobantePago)

		// Logs (solo admin)
		logs := api.Group("/logs")
		logs.Use(middleware.RoleMiddleware("admin"))
		{
			logs.GET("", logHandler.GetLogs)
		}
	}

	// Iniciar servidor
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
