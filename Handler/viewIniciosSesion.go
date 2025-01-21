package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewIniciosSesion(c *gin.Context) {
	// Consulta para obtener los inicios de sesión de los usuarios
	query := `
		SELECT 
			logLogin.fecha_hora, 
			CONCAT(users.nombre, ' ', users.apellido) AS nombre_completo, 
			users.codigo_empleado
		FROM 
			logLogin
		INNER JOIN 
			users 
		ON 
			users.id = logLogin.id_usuario
	`

	rows, err := config.DBTurso.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los inicios de sesión", "details": err.Error()})
		return
	}
	defer rows.Close()

	// Crear un slice para almacenar los datos de los inicios de sesión
	var iniciosSesion []models.InicioSesion

	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		var inicioSesion models.InicioSesion
		if err := rows.Scan(&inicioSesion.FechaHora, &inicioSesion.NombreCompleto, &inicioSesion.CodigoEmpleado); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los inicios de sesión", "details": err.Error()})
			return
		}
		iniciosSesion = append(iniciosSesion, inicioSesion)
	}

	// Verificar si no hay registros
	if len(iniciosSesion) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay registros de inicios de sesión"})
		return
	}

	// Respuesta exitosa con los inicios de sesión
	c.JSON(http.StatusOK, gin.H{
		"message": "Inicios de sesión obtenidos exitosamente",
		"inicios_sesion": iniciosSesion,
	})
}
