package main

import (
	"application-preparing-room/config"
	"application-preparing-room/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configurar Gin
	gin.SetMode(gin.ReleaseMode)

	// Conectar a la base de datos
	if err := config.ConnectDB(); err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Registrar rutas
	router := routes.RegisterRoutes()

	// Iniciar el servidor
	PORT := config.GetPort()
	log.Printf("Servidor corriendo en el puerto %s", PORT)
	router.Run(":" + PORT)
}
