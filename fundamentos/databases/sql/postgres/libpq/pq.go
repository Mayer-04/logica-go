package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

func main() {
	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open(driverName, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Verificar la conexión
	if error := db.Ping(); error != nil {
		log.Fatal("Error al conectarnos a la base de datos")
	}

	var age int
	// Obtener varias filas
	_, err = db.Query("SELECT name FROM users WHERE age = $1", age)

	if err != nil {
		log.Fatal(err)
	}
}
