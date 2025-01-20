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
	r.GET("/viewRoom",Handler.ViewRoom) //ruta para visualizar todas las habitaciones
	r.GET("/viewUser",Handler.ViewUsers) //ruta para visualizar a todos los usuarios
	r.GET("/viewAllTask",Handler.ViewAllTasks) //ruta para visualizar todas las tareas
	r.POST("/register", Handler.Register) // Ruta para registrar un usuario
	r.POST("/login",Handler.Login)
	r.POST("/createRoom",Handler.CreateRoom) //ruta para crear habitaciones
	r.POST("/assignTask",Handler.AssignTask) //ruta para asignar la tarea de limpieza de habitacion
	r.POST("/updateStatusTask",Handler.UpdateStatusTask) //ruta para actualizar la columna status
	r.DELETE("/deleteUser/:codigo_empleado", Handler.DeleteUser)
	r.PUT("/updateToAdmin/:codigo_empleado", Handler.UpdateToAdmin)
	r.DELETE("/deleteRoom/:id", Handler.DeleteRoom) // Ruta para eliminar una habitación
	r.PUT("/updateRoom/:id",Handler.UpdateRoom) //ruta para actualizar una habitacion


	return r
}
