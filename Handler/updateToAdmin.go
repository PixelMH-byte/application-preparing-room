package Handler

import (
	"application-preparing-room/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"log"
)

func UpdateToAdmin(c *gin.Context) {
	// Capturar el parámetro del código del empleado desde la URL
	codigoEmpleado := c.Param("codigo_empleado")

	// Log para depuración
	log.Println("Código empleado recibido desde la URL:", codigoEmpleado)

	var exists bool
	queryCheck := "SELECT EXISTS(SELECT 1 FROM users WHERE codigo_empleado = ?)"
	err := config.DBTurso.QueryRow(queryCheck, codigoEmpleado).Scan(&exists)
	if err != nil || !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "El usuario no existe"})
		return
	}

	// Actualizar el campo profile a 1 (true)
	queryUpdate := "UPDATE users SET profile = 1 WHERE codigo_empleado = ?"
	_, err = config.DBTurso.Exec(queryUpdate, codigoEmpleado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario a administrador", "details": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario actualizado a administrador exitosamente",
		"codigo_empleado": codigoEmpleado,
	})
}

