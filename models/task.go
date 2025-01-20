package models

// TaskRequest estructura para capturar los datos necesarios para asignar una tarea
type TaskRequest struct {
	CodigoEmpleado string `json:"codigo_empleado" binding:"required"` // Código del empleado asignado
	TaskName       string `json:"task_name" binding:"required"`       // Nombre de la tarea
	NumHabitacion  int    `json:"num_habitacion" binding:"required"`  // Número de la habitación
	Planta         int    `json:"planta" binding:"required"`          // Planta de la habitación
}

// Task estructura para representar la tabla taskRoom
type Task struct {
	ID             int    `json:"id"`
	TaskName       string `json:"task_name"`
	Status         string `json:"status"`
	RoomID         int    `json:"room_id"`
	CodigoEmpleado string `json:"codigo_empleado"`
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
}
