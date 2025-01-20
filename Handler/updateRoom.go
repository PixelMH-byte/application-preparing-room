package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRoom(c *gin.Context) {
	// Capturar el parámetro del ID de la habitación desde la URL
	id := c.Param("id")

	// Verificar que el ID sea un número válido
	roomID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe ser un número válido"})
		return
	}

	// Crear una instancia del struct Room para capturar los datos enviados
	var roomUpdate models.RoomUpdateRequest

	// Bind JSON del cuerpo de la solicitud al struct
	if err := c.ShouldBindJSON(&roomUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si la habitación existe
	var exists bool
	queryCheck := "SELECT EXISTS(SELECT 1 FROM room WHERE id = ?)"
	err = config.DBTurso.QueryRow(queryCheck, roomID).Scan(&exists)
	if err != nil || !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "La habitación no existe"})
		return
	}

	// Actualizar los datos de la habitación
	queryUpdate := "UPDATE room SET planta = ?, tamaño = ?, num_habitacion = ? WHERE id = ?"
	_, err = config.DBTurso.Exec(queryUpdate, roomUpdate.Planta, roomUpdate.Tamaño, roomUpdate.NumHabitacion, roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la habitación", "details": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Habitación actualizada exitosamente",
		"room": gin.H{
			"id":             roomID,
			"planta":         roomUpdate.Planta,
			"tamano":         roomUpdate.Tamaño,
			"num_habitacion": roomUpdate.NumHabitacion,
		},
	})
}
