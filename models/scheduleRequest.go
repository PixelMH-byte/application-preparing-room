package models

// ScheduleRequest estructura para capturar los datos necesarios para asignar un horario
type ScheduleRequest struct {
	Turno          string `json:"turno" binding:"required"`             // Turno (mañana, tarde, noche)
	NumHoras       int    `json:"num_horas" binding:"required"`         // Número de horas asignadas
	CodigoEmpleado string `json:"codigo_empleado" binding:"required"`   // Código del empleado
	StartDayTime   string `json:"start_day_time" binding:"required"`    // Fecha y hora de inicio
	FinishDayTime  string `json:"finish_day_time" binding:"required"`   // Fecha y hora de fin
}
