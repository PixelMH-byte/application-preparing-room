package Handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.String(http.StatusOK, "¡Bienvenido a la API del proyecto de localizacion de cablescom!")
}