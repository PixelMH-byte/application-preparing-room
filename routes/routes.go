package routes

import (
	"application-preparing-room/Handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes configura las rutas y middlewares del servidor
func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // Permitir todas las orígenes
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * 3600, // Tiempo máximo de preflight (12 horas)
	}))

	// Rutas de la API
	r.GET("/", Handler.Welcome)          // Ruta de prueba para verificar el servidor
	r.POST("/register", Handler.Register) // Ruta para registrar un usuario
	r.POST("/login",Handler.Login)
	return r
}
