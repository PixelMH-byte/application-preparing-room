package models

// InicioSesion estructura para representar los datos de los inicios de sesión
type InicioSesion struct {
	FechaHora      string `json:"fecha_hora"`
	NombreCompleto string `json:"nombre_completo"`
	CodigoEmpleado string `json:"codigo_empleado"`
}
