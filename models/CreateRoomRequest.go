package models

type CreateRoomRequest struct {
	CodigoEmpleado string `json:"codigo_empleado" binding:"required"`
	Planta         int    `json:"planta" binding:"required"`
	Tamaño         string `json:"tamano" binding:"required"`
	Num_Habitacion int    `json:"num_habitacion" binding:"required"`
}
