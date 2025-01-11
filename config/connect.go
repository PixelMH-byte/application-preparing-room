package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Asegúrate de importar este paquete
	_ "github.com/tursodatabase/libsql-client-go/libsql" // Driver libSQL
)


var DBTurso *sql.DB // Declarar la variable global para la conexión a Turso

func init() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
}

func ConnectDB() error {
	// Obtener las credenciales de Turso desde las variables de entorno
	tursoURL := os.Getenv("TURSO_URL")
	tursoAPIToken := os.Getenv("TURSO_API_TOKEN")

	// Validar que las variables no estén vacías
	if tursoURL == "" || tursoAPIToken == "" {
		return fmt.Errorf("error: TURSO_URL o TURSO_API_TOKEN no están configurados")
	}

	// Construir el DSN (Data Source Name) para la conexión
	dsn := fmt.Sprintf("%s?authToken=%s", tursoURL, tursoAPIToken)

	// Abrir la conexión a la base de datos usando el driver correcto
	db, err := sql.Open("libsql", dsn)
	if err != nil {
		return fmt.Errorf("error creando la conexión a Turso: %v", err)
	}

	// Verificar la conexión
	if err = db.Ping(); err != nil {
		return fmt.Errorf("error verificando la conexión a Turso: %v", err)
	}

	DBTurso = db
	log.Println("Conexión exitosa a la base de datos de Turso")
	return nil
}

func GetPort() string {
	// Obtener el puerto de la variable de entorno
	PORT := os.Getenv("HTTP_PLATFORM_PORT")
	if PORT == "" {
		PORT = "8100" // Valor predeterminado si no se proporciona un puerto
	}
	return PORT
}
