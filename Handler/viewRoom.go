package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewRoom(c *gin.Context) {
	// Consulta SQL para obtener los datos de la tabla room
	query := "SELECT id, planta, tamaño,num_habitacion FROM room"

	// Slice para almacenar los resultados
	var rooms []models.Room

	// Ejecutar la consulta
	rows, err := config.DBTurso.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las habitaciones", "details": err.Error()})
		return
	}
	defer rows.Close()

	// Recorrer los resultados
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.ID, &room.Planta, &room.Tamaño, &room.Num_habitacion); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer los datos de las habitaciones", "details": err.Error()})
			return
		}
		rooms = append(rooms, room)
	}

	// Verificar si no hay resultados
	if len(rooms) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No hay habitaciones disponibles"})
		return
	}

	// Respuesta exitosa con los datos
	c.JSON(http.StatusOK, gin.H{
		"rooms": rooms,
	})
}
