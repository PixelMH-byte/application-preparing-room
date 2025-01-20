package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewTaskUser(c *gin.Context) {
	// Capturar el parámetro del código del empleado desde la URL
	codigoEmpleado := c.Param("codigo_empleado")

	// Consultar las tareas asignadas al empleado con el código proporcionado
	query := `
		SELECT 
			taskRoom.task_name, 
			taskRoom.status, 
			room.planta, 
			room.num_habitacion, 
			users.codigo_empleado
		FROM 
			taskRoom
		INNER JOIN 
			users 
		ON 
			users.codigo_empleado = taskRoom.codigo_empleado
		INNER JOIN 
			room 
		ON 
			room.id = taskRoom.roomid
		WHERE 
			users.codigo_empleado = ?
	`
	rows, err := config.DBTurso.Query(query, codigoEmpleado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las tareas del usuario", "details": err.Error()})
		return
	}
	defer rows.Close()

	// Crear un slice para almacenar las tareas
	var tasks []models.TaskDetails

	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		var task models.TaskDetails
		if err := rows.Scan(&task.TaskName, &task.Status, &task.Planta, &task.NumHabitacion, &task.CodigoEmpleado); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar las tareas del usuario", "details": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	// Verificar si no hay tareas asignadas
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay tareas asignadas a este usuario"})
		return
	}

	// Respuesta exitosa con las tareas del usuario
	c.JSON(http.StatusOK, gin.H{
		"message": "Tareas obtenidas exitosamente",
		"tasks":   tasks,
	})
}
