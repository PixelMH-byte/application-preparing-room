package models

// Room estructura que representa una habitación en la tabla room
type Room struct {
	ID     int    `json:"id"`
	Planta int    `json:"planta"`
	Tamaño string `json:"tamano"`
	Num_habitacion int    `json:"num_habitacion"`

	
}
