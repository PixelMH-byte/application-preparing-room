package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewUsers(c *gin.Context) {
	// Consulta SQL para obtener los datos de la tabla users
	query := "SELECT nombre, apellido, codigo_empleado, profile FROM users"

	// Slice para almacenar los resultados
	var users []models.User

	// Ejecutar la consulta
	rows, err := config.DBTurso.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios", "details": err.Error()})
		return
	}
	defer rows.Close()

	// Recorrer los resultados
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Nombre, &user.Apellido, &user.CodigoEmpleado, &user.Profile); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer los datos de los usuarios", "details": err.Error()})
			return
		}
		users = append(users, user)
	}

	// Verificar si no hay resultados
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No hay usuarios disponibles"})
		return
	}

	// Respuesta exitosa con los datos
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
