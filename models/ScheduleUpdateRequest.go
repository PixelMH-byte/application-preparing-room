package models

// ScheduleUpdateRequest estructura para capturar los datos necesarios para actualizar un horario
type ScheduleUpdateRequest struct {
	Turno          string `json:"turno" binding:"required"`             // Turno (mañana, tarde, noche)
	NumHoras       int    `json:"num_horas" binding:"required"`         // Número de horas asignadas
	StartDayTime   string `json:"start_day_time" binding:"required"`    // Fecha y hora de inicio
	FinishDayTime  string `json:"finish_day_time" binding:"required"`   // Fecha y hora de fin
}
