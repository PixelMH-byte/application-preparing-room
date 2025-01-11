package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	// Vincular JSON al struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Validar y asignar valores predeterminados
	if !user.Profile {
		user.Profile = false
	}

	// Generar código de empleado
	rand.Seed(time.Now().UnixNano())
	user.CodigoEmpleado = "PIXEL" + randomFourDigitCode()

	// Validar datos generados
	fmt.Printf("Insertando usuario: Nombre=%s, Apellido=%s, Profile=%v, CodigoEmpleado=%s\n",
		user.Nombre, user.Apellido, user.Profile, user.CodigoEmpleado)

	// Consulta de inserción
	queryInsert := "INSERT INTO users (nombre, apellido, profile, codigo_empleado) VALUES (?, ?, ?, ?)"
	result, err := config.DBTurso.Exec(queryInsert, user.Nombre, user.Apellido, user.Profile, user.CodigoEmpleado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar el usuario", "details": err.Error()})
		return
	}

	// Obtener el ID del usuario recién insertado
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el ID del usuario", "details": err.Error()})
		return
	}
	user.ID = int(id)

	// Consultar fecha_creacion
	querySelect := "SELECT fecha_creacion FROM users WHERE id = ?"
	row := config.DBTurso.QueryRow(querySelect, user.ID)
	if err := row.Scan(&user.FechaCreacion); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al recuperar la fecha de creación", "details": err.Error()})
		return
	}

	// Respuesta al cliente
	c.JSON(http.StatusCreated, gin.H{
		"message": "Usuario registrado exitosamente",
		"user":    user,
	})
}

func randomFourDigitCode() string {
	return fmt.Sprintf("%04d", rand.Intn(10000))
}
