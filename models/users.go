package models

type User struct {
	ID             int    `json:"id"`
	Nombre         string `json:"nombre"`
	Apellido       string `json:"apellido"`
	Profile        bool   `json:"profile"`
	FechaCreacion  string `json:"fecha_creacion"`
	CodigoEmpleado string `json:"codigo_empleado"`
}
