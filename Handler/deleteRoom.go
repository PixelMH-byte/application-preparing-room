package Handler

import (
	"application-preparing-room/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteRoom(c *gin.Context) {
	// Capturar el parámetro del ID de la habitación desde la URL
	id := c.Param("id")

	// Verificar si la habitación existe
	var exists bool
	queryCheck := "SELECT EXISTS(SELECT 1 FROM room WHERE id = ?)"
	err := config.DBTurso.QueryRow(queryCheck, id).Scan(&exists)
	if err != nil || !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "La habitación no existe"})
		return
	}

	// Eliminar la habitación de la tabla room
	queryDelete := "DELETE FROM room WHERE id = ?"
	_, err = config.DBTurso.Exec(queryDelete, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar la habitación", "details": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Habitación eliminada exitosamente",
		"id":      id,
	})
}
