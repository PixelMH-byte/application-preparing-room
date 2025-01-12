package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssignTask(c *gin.Context) {
	// Crear una instancia del struct TaskRequest para capturar los datos de la solicitud
	var taskRequest models.TaskRequest

	// Bind JSON del cuerpo de la solicitud al struct
	if err := c.ShouldBindJSON(&taskRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si el empleado existe en la tabla users y si tiene permisos (profile = true)
	var profile bool
	queryProfile := "SELECT profile FROM users WHERE codigo_empleado = ?"
	err := config.DBTurso.QueryRow(queryProfile, taskRequest.CodigoEmpleado).Scan(&profile)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "El código de empleado no existe o no tiene permisos"})
		return
	}

	// Validar si el usuario tiene perfil de administrador
	if !profile {
		c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para asignar tareas"})
		return
	}

	// Buscar el ID de la habitación usando el número de habitación y la planta
	var roomID int
	queryRoom := "SELECT id FROM room WHERE planta = ? AND num_habitacion = ?"
	err = config.DBTurso.QueryRow(queryRoom, taskRequest.Planta, taskRequest.NumHabitacion).Scan(&roomID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontró una habitación con los datos proporcionados"})
		return
	}


	// Insertar la tarea en la tabla taskRoom
	queryInsert := "INSERT INTO taskRoom (task_name, status, roomid, codigo_empleado) VALUES (?, ?, ?, ?)"
	_, err = config.DBTurso.Exec(queryInsert, taskRequest.TaskName, "pending", roomID, taskRequest.CodigoEmpleado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la tarea", "details": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusCreated, gin.H{
		"message": "Tarea asignada exitosamente",
		"task": gin.H{
			"task_name":       taskRequest.TaskName,
			"status":          "pending",
			"room_id":         roomID,
			"codigo_empleado": taskRequest.CodigoEmpleado,
		},
	})
}
