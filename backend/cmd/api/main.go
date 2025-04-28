package main

import (
	"backend/cmd/api/routes"
	"backend/internal/config/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
	database.Connect()
	defer database.DB.Close()

	// Ejecutar migraciones
	database.Migrate()

	// Configurar el router de Gin
	router := gin.Default()

	// Configurar rutas
	routes.SetupUserRoutes(database.DB, router)

	// Configuración del servidor
	port := ":8080"
	log.Printf("Servidor iniciado en http://localhost%s", port)

	// Iniciar el servidor
	if err := router.Run(port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
