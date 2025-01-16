package Handler

import (
	"application-preparing-room/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	// Capturar el parámetro del código del empleado desde la URL
	codigoEmpleado := c.Param("codigo_empleado")

	// Verificar si el usuario existe
	var userExists bool
	queryCheck := "SELECT EXISTS(SELECT 1 FROM users WHERE codigo_empleado = ?)"
	err := config.DBTurso.QueryRow(queryCheck, codigoEmpleado).Scan(&userExists)
	if err != nil || !userExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "El usuario no existe"})
		return
	}

	// Eliminar el usuario de la tabla
	queryDelete := "DELETE FROM users WHERE codigo_empleado = ?"
	_, err = config.DBTurso.Exec(queryDelete, codigoEmpleado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario", "details": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario eliminado exitosamente",
		"codigo_empleado": codigoEmpleado,
	})
}
