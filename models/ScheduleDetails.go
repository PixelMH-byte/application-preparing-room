package models

// ScheduleDetails estructura para representar los detalles de un horario
type ScheduleDetails struct {
	ID             int    `json:"id"`
	CodigoEmpleado string `json:"codigo_empleado"`
	NombreCompleto string `json:"nombre_completo"`
	Turno          string `json:"turno"`
	NumHoras       int    `json:"num_horas"`
	StartDayTime   string `json:"start_day_time"`
	FinishDayTime  string `json:"finish_day_time"`
}
