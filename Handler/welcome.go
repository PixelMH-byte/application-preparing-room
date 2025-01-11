package Handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome es una función de prueba para la ruta "/"
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Preparing Room API!",
	})
}
