package models

// UpdateTaskRequest estructura para capturar los datos necesarios para actualizar el estado de una tarea
type UpdateTaskRequest struct {
	CodigoEmpleado string `json:"codigo_empleado" binding:"required"` // Código del empleado que intenta actualizar
	TaskID         int    `json:"task_id" binding:"required"`         // ID de la tarea a actualizar
	NewStatus      string `json:"new_status" binding:"required"`      // Nuevo estado de la tarea
}
