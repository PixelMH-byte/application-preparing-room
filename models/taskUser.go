package models

// TaskDetails estructura que representa los detalles de una tarea
type TaskDetails struct {
	ID             int    `json:"id"`
	TaskName       string `json:"task_name"`
	Status         string `json:"status"`
	Planta         int    `json:"planta"`
	NumHabitacion  int    `json:"num_habitacion"`
	CodigoEmpleado string `json:"codigo_empleado"`
}
