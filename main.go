package main

import (
    "log"
    "net/http"
    "application-preparing-room/config"  // Usa el nombre real de tu módulo

    "github.com/gin-gonic/gin"
)

func main() {
    // Establecer el modo de Gin en 'release' para producción
    gin.SetMode(gin.ReleaseMode)

    // Conectar a la base de datos
    if err := config.ConnectDB(); err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    // Crear un nuevo router Gin
    r := gin.Default()

    // Ruta de prueba
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    // Aquí puedes agregar más rutas o importar tu archivo de rutas
    // routes.RegisterRoutes(r)

    // Obtener el puerto desde la configuración
    PORT := config.GetPort()

    // Iniciar el servidor con Gin
    log.Printf("Server running on port %s", PORT)
    log.Fatal(r.Run(":" + PORT))
}