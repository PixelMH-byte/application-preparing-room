package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Estructura para capturar los datos del login
	var loginRequest struct {
		Nombre         string `json:"nombre"`
		CodigoEmpleado string `json:"codigo_empleado"`
	}

	// Bind JSON del cuerpo de la solicitud al struct
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar las credenciales en la tabla users
	var user models.User
	querySelect := "SELECT id, nombre FROM users WHERE nombre = ? AND codigo_empleado = ?"
	err := config.DBTurso.QueryRow(querySelect, loginRequest.Nombre, loginRequest.CodigoEmpleado).
		Scan(&user.ID, &user.Nombre)

	if err == sql.ErrNoRows {
		// No se encontró un usuario con las credenciales proporcionadas
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	} else if err != nil {
		// Otro error durante la consulta
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al verificar las credenciales", "details": err.Error()})
		return
	}

	// Insertar en la tabla logLogin
	queryInsert := "INSERT INTO logLogin (id_usuario) VALUES (?)"
	_, err = config.DBTurso.Exec(queryInsert, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el login", "details": err.Error()})
		return
	}

	// Respuesta al cliente
	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"user":    user.Nombre,
	})
}
