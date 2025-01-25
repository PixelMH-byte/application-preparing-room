package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateSchedulesUser(c *gin.Context) {
	// Capturar el parámetro del código del empleado desde la URL
	codigoEmpleado := c.Param("codigo_empleado")

	// Verificar si el código del empleado existe en la tabla schedules_users
	var idUsuario int
	queryUser := "SELECT id_usuario FROM schedules_users WHERE codigo_empleado = ?"
	err := config.DBTurso.QueryRow(queryUser, codigoEmpleado).Scan(&idUsuario)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "El código de empleado no existe o no tiene un horario asignado"})
		return
	}

	// Crear una instancia del struct ScheduleUpdateRequest para capturar los datos de la solicitud
	var scheduleUpdate models.ScheduleUpdateRequest

	// Bind JSON del cuerpo de la solicitud al struct
	if err := c.ShouldBindJSON(&scheduleUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Actualizar los campos del horario
	queryUpdate := `
		UPDATE schedules_users 
		SET turno = ?, num_horas = ?, start_day_time = ?, finish_day_time = ? 
		WHERE codigo_empleado = ?
	`
	_, err = config.DBTurso.Exec(queryUpdate, scheduleUpdate.Turno, scheduleUpdate.NumHoras, scheduleUpdate.StartDayTime, scheduleUpdate.FinishDayTime, codigoEmpleado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el horario", "details": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Horario actualizado exitosamente",
		"schedule": gin.H{
			"codigo_empleado":  codigoEmpleado,
			"turno":           scheduleUpdate.Turno,
			"num_horas":       scheduleUpdate.NumHoras,
			"start_day_time":  scheduleUpdate.StartDayTime,
			"finish_day_time": scheduleUpdate.FinishDayTime,
		},
	})
}
