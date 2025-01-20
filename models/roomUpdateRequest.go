package models

// RoomUpdateRequest estructura para capturar los datos necesarios para actualizar una habitación
type RoomUpdateRequest struct {
	Planta         int    `json:"planta" binding:"required"`
	Tamaño         string `json:"tamano" binding:"required"`
	NumHabitacion  int    `json:"num_habitacion" binding:"required"`
}
