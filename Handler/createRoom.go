package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRoom(c *gin.Context) {
	// Crear una instancia del struct para capturar los datos de la solicitud
	var roomRequest models.CreateRoomRequest

	// Bind JSON del cuerpo de la solicitud al struct
	if err := c.ShouldBindJSON(&roomRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si el usuario tiene perfil de administrador
	var profile bool
	queryProfile := "SELECT profile FROM users WHERE codigo_empleado = ?"
	err := config.DBTurso.QueryRow(queryProfile, roomRequest.CodigoEmpleado).Scan(&profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar el perfil del usuario", "details": err.Error()})
		return
	}

	// Validar si el usuario tiene perfil de administrador
	if !profile {
		c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para crear habitaciones"})
		return
	}

	// Insertar la habitación en la tabla room
	queryInsert := "INSERT INTO room (planta, tamaño, num_habitacion) VALUES (?, ?, ?)" // Aquí se corrige la cantidad de placeholders
	_, err = config.DBTurso.Exec(queryInsert, roomRequest.Planta, roomRequest.Tamaño, roomRequest.Num_Habitacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la habitación", "details": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusCreated, gin.H{
		"message": "Habitación creada exitosamente",
		"room": gin.H{
			"planta":         roomRequest.Planta,
			"tamano":         roomRequest.Tamaño,
			"num_habitacion": roomRequest.Num_Habitacion,
		},
	})
}
