package models

type LogLogin struct {
	ID        int    `json:"id"`
	FechaHora string `json:"fecha_hora"`
	IDUsuario int    `json:"id_usuario"`
	User      string `json:"user"`
}
