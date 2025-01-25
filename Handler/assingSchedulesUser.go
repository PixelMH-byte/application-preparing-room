package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssignSchedulesUser(c *gin.Context) {
	// Crear una instancia del struct ScheduleRequest para capturar los datos de la solicitud
	var scheduleRequest models.ScheduleRequest

	// Bind JSON del cuerpo de la solicitud al struct
	if err := c.ShouldBindJSON(&scheduleRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si el código del empleado existe en la tabla users y obtener id_usuario
	var idUsuario int
	queryUser := "SELECT id FROM users WHERE codigo_empleado = ?"
	err := config.DBTurso.QueryRow(queryUser, scheduleRequest.CodigoEmpleado).Scan(&idUsuario)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "El código de empleado no existe"})
		return
	}

	// Insertar el horario en la tabla schedules_users
	queryInsert := `
		INSERT INTO schedules_users (turno, num_horas, id_usuario, codigo_empleado, start_day_time, finish_day_time) 
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err = config.DBTurso.Exec(queryInsert, scheduleRequest.Turno, scheduleRequest.NumHoras, idUsuario, scheduleRequest.CodigoEmpleado, scheduleRequest.StartDayTime, scheduleRequest.FinishDayTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al asignar el horario", "details": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusCreated, gin.H{
		"message": "Horario asignado exitosamente",
		"schedule": gin.H{
			"turno":            scheduleRequest.Turno,
			"num_horas":        scheduleRequest.NumHoras,
			"codigo_empleado":  scheduleRequest.CodigoEmpleado,
			"start_day_time":   scheduleRequest.StartDayTime,
			"finish_day_time":  scheduleRequest.FinishDayTime,
		},
	})
}
