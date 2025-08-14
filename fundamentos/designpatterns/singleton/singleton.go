package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

/*
* Patrón de diseño Singleton
El patrón Singleton garantiza que solo exista una única instancia de una clase (o estructura en Go)
a lo largo del ciclo de vida de una aplicación.

- El propósito principal de este patrón es controlar el acceso a recursos compartidos,
como una conexión a la base de datos o un archivo.
- Proporciona un único punto de acceso global a la instancia.
- Es importante tener en cuenta que este patrón puede violar el principio de responsabilidad única,
ya que combina la gestión de la instancia con su funcionalidad.
- En Go, el uso de `sync.Once` garantiza que la inicialización de la instancia sea segura incluso en entornos
concurrentes, ejecutando la función de inicialización una única vez.
*/

// database es una estructura que encapsula la conexión a la base de datos.
type Database struct {
	Conn *sql.DB
}

// Las variables globales pueden ser peligrosas ya que cualquier parte del código puede modificar su contenido.
var (
	// dbInstance almacena la única instancia de la estructura database.
	dbInstance *Database
	// once se utiliza para asegurarse de que la instancia de la base de datos solo se inicialice una vez.
	once sync.Once
)

// connectionDB configura y establece una conexión con la base de datos PostgreSQL.
// Si la conexión es exitosa, la instancia se almacena en la variable dbInstance.
// En caso de error al abrir la conexión o al verificarla (ping), el programa se termina con un mensaje de error.
func connectionDB() *Database {
	// Configuración de la cadena de conexión a PostgreSQL.
	connStr := "user=youruser dbname=yourdb password=yourpassword host=localhost sslmode=disable"
	// Intenta abrir una conexión con la base de datos.
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al abrir la conexión a la base de datos: %v", err)
	}

	// Verifica si la conexión con la base de datos es válida y está activa.
	// db.PingContext()
	if err := db.Ping(); err != nil {
		log.Fatalf("Error al hacer ping a la base de datos: %v", err)
	}

	// Asigna la conexión a la instancia singleton.
	return &Database{Conn: db}
}

// GetDBInstance devuelve la instancia única de la conexión a la base de datos.
// Utiliza `sync.Once` para asegurar que la instancia solo se cree una vez, incluso en un entorno concurrente.
func GetDBInstance() *Database {
	// La función connectionDB solo se ejecuta una vez, gracias a `sync.Once`.
	once.Do(func() {
		dbInstance = connectionDB()
	})
	return dbInstance
}

func main() {
	// Obtener la instancia única de la conexión a la base de datos.
	db1 := GetDBInstance()
	db2 := GetDBInstance()

	// Verificar si ambas variables apuntan a la misma instancia.
	if db1 == db2 {
		fmt.Println("db1 y db2 son la misma instancia.")
	}
}
