package Handler

import (
	"application-preparing-room/config"
	"application-preparing-room/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStatusTask(c *gin.Context) {
	// Crear una instancia del struct UpdateTaskRequest para capturar los datos de la solicitud
	var updateRequest models.UpdateTaskRequest

	// Bind JSON del cuerpo de la solicitud al struct
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Verificar si el estado proporcionado es válido
	validStatuses := map[string]bool{
		"pending": true,
		"progress": true,
		"finish": true,
	}
	if !validStatuses[updateRequest.NewStatus] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estado no válido"})
		return
	}

	// Verificar si el empleado existe y obtener su perfil
	var profile bool
	queryProfile := "SELECT profile FROM users WHERE codigo_empleado = ?"
	err := config.DBTurso.QueryRow(queryProfile, updateRequest.CodigoEmpleado).Scan(&profile)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "El código de empleado no existe"})
		return
	}

	// Validar si el nuevo estado es "pending" y el usuario no tiene perfil de administrador
	if updateRequest.NewStatus == "pending" && !profile {
		c.JSON(http.StatusForbidden, gin.H{"error": "Solo los administradores pueden volver a establecer el estado a 'pending'"})
		return
	}

	// Validar si el usuario sin perfil de administrador intenta cambiar el estado a progress o finish
	if updateRequest.NewStatus != "pending" && !profile {
		if updateRequest.NewStatus != "progress" && updateRequest.NewStatus != "finish" {
			c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permisos para realizar esta acción"})
			return
		}
	}

	// Actualizar el estado de la tarea en la tabla taskRoom
	queryUpdate := "UPDATE taskRoom SET status = ? WHERE id = ?"
	result, err := config.DBTurso.Exec(queryUpdate, updateRequest.NewStatus, updateRequest.TaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el estado de la tarea", "details": err.Error()})
		return
	}

	// Verificar si se actualizó alguna fila
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo verificar la actualización", "details": err.Error()})
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontró la tarea especificada"})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Estado de la tarea actualizado correctamente",
		"task": gin.H{
			"task_id":      updateRequest.TaskID,
			"new_status":   updateRequest.NewStatus,
			"updated_by":   updateRequest.CodigoEmpleado,
		},
	})
}
