package models

// UpdateToAdminRequest estructura para capturar los datos necesarios para actualizar el perfil a administrador
type UpdateToAdminRequest struct {
	CodigoEmpleado string `json:"codigo_empleado" binding:"required"` // Código del empleado que será actualizado
}
