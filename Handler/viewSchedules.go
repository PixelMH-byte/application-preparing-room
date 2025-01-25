package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewSchedules(c *gin.Context) {
	// Consulta para obtener todos los horarios junto con información de los trabajadores
	query := `
		SELECT 
			schedules_users.id, 
			users.codigo_empleado, 
			CONCAT(users.nombre, ' ', users.apellido) AS nombre_completo,
			schedules_users.turno, 
			schedules_users.num_horas, 
			schedules_users.start_day_time, 
			schedules_users.finish_day_time
		FROM 
			schedules_users
		INNER JOIN 
			users 
		ON 
			users.id = schedules_users.id_usuario
	`

	rows, err := config.DBTurso.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los horarios", "details": err.Error()})
		return
	}
	defer rows.Close()

	// Crear un slice para almacenar los horarios
	var schedules []models.ScheduleDetails

	// Iterar sobre los resultados de la consulta
	for rows.Next() {
		var schedule models.ScheduleDetails
		if err := rows.Scan(&schedule.ID, &schedule.CodigoEmpleado, &schedule.NombreCompleto, &schedule.Turno, &schedule.NumHoras, &schedule.StartDayTime, &schedule.FinishDayTime); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los horarios", "details": err.Error()})
			return
		}
		schedules = append(schedules, schedule)
	}

	// Verificar si no hay horarios
	if len(schedules) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No hay horarios registrados"})
		return
	}

	// Respuesta exitosa con los horarios
	c.JSON(http.StatusOK, gin.H{
		"message":   "Horarios obtenidos exitosamente",
		"schedules": schedules,
	})
}
