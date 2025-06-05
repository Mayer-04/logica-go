package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool" // Controlador de base de datos
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// "postgres://user:password@localhost:5432/dbname
	// puerto y nombre de la base de datos pueden ser opcionales.
	dbpool, err := pgxpool.New(ctx, "postgres://mayer:123456@localhost?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer dbpool.Close()

	// Ejecuta una sentencia SQL vacía.
	// Si la setencia no devuelve un error, la conexión a la base de datos fue exitosa.
	// Si no usas Ping(), no sabrás si la base de datos está disponible hasta que intentes hacer una consulta.
	if err := dbpool.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("Conexión exitosa... ✅")
}
