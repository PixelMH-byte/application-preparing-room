package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewAllTasks(c *gin.Context) {
	// Consultar todas las tareas de la tabla taskRoom junto con los datos del usuario asignado
	query := `
		SELECT taskRoom.id, taskRoom.task_name, taskRoom.status, taskRoom.roomid, 
		       users.codigo_empleado, users.nombre, users.apellido
		FROM taskRoom 
		INNER JOIN users ON users.codigo_empleado = taskRoom.codigo_empleado
	`
	rows, err := config.DBTurso.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las tareas", "details": err.Error()})
		return
	}
	defer rows.Close()

	// Crear un slice para almacenar las tareas
	var tasks []models.Task

	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.TaskName, &task.Status, &task.RoomID, &task.CodigoEmpleado, &task.Nombre, &task.Apellido); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar las tareas", "details": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	// Verificar si no hay tareas
	if len(tasks) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay tareas registradas"})
		return
	}

	// Respuesta exitosa con las tareas
	c.JSON(http.StatusOK, gin.H{
		"message": "Tareas obtenidas exitosamente",
		"tasks":   tasks,
	})
}
