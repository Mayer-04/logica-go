package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

/*
* Servidor web estático en Go
Este programa crea un servidor web que sirve archivos estáticos desde una carpeta específica (public).

* Paquetes utilizados en el código:
- El paquete "path/filepath" proporciona funciones para trabajar con rutas de archivos y directorios.
- El paquete "os" proporciona funciones para interactuar con el sistema operativo,
como obtener el directorio de trabajo actual.
- El paquete "slog" se utiliza para registrar eventos y mensajes informativos en el sistema de registro.
*/

func main() {
	// Configurando un nuevo registrador de mensajes.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	// Obtiene la ruta completa a la carpeta 'public', donde se encuentran los archivos estáticos.
	publicPath := getPublicPath()
	// Crea un servidor de archivos estáticos que servirá los archivos desde la carpeta 'public'.
	fs := http.FileServer(http.Dir(publicPath))
	// Configura el manejador de solicitudes HTTP para servir archivos estáticos desde la ruta raíz (/).
	http.Handle("/", fs)
	// Definimos el puerto en el que el servidor HTTP escuchará las solicitudes.
	addr := ":5000"

	// Configuramos el servidor HTTP con tiempos de espera personalizados para lectura y escritura.
	// Recomendado para apliaciones de producción.
	server := &http.Server{
		Addr: addr,
		// Si el cliente no envía datos en un plazo de tiempo definido, el servidor cierra la conexión.
		// El cliente puede volver a hacer una solicitud después de que el servidor cierre una conexión.
		ReadTimeout: 5 * time.Second,
		// Si el cliente no pudo recibir la respuesta completa del servidor, el servidor cierra la conexión.
		WriteTimeout: 5 * time.Second,
		Handler:      fs,
	}

	// Registra un mensaje en la consola utilizando slog cuando el servidor se inicia.
	logger.Info("Iniciando servidor", "puerto", addr)

	// Iniciamos el servidor web en el puerto especificado.
	if err := server.ListenAndServe(); err != nil {
		// Imprime el error y termina el programa si no se puede iniciar el servidor.
		logger.Error("fallo al iniciar el servidor", "error", err)
		os.Exit(1)
	}
}

// getCurrentDirectory obtiene el directorio de trabajo actual.
func getCurrentDirectory() (string, error) {
	// Obtener el directorio de trabajo actual.
	dir, err := os.Getwd()
	if err != nil {
		// Devuelve un error envuelto con un mensaje explicativo si no se puede obtener el directorio.
		return "", fmt.Errorf("error obteniendo el directorio: %w", err)
	}
	// Devuelve la ruta del directorio actual si no hubo errores.
	return dir, nil
}

// getPublicPath construye la ruta completa a la carpeta 'public' donde se almacenan los archivos estáticos.
func getPublicPath() string {
	// Obtiene el directorio de trabajo actual.
	dir, err := getCurrentDirectory()
	if err != nil {
		// Imprime el error si no se puede obtener el directorio de trabajo actual.
		fmt.Println(err)
		// Termina el programa con un estado de error genérico.
		// El código 1 indica una falla por una razón que no está claramente definida.
		os.Exit(1)
	}

	// Define una lista de directorios que componen la ruta relativa a la carpeta 'public'.
	directorys := [...]string{"fundamentos", "server", "serverFile", "public"}
	for _, directory := range directorys {
		// Acumula la ruta completa añadiendo cada directorio de la lista.
		dir = filepath.Join(dir, directory)
	}
	// Devuelve la ruta completa a la carpeta 'public'.
	return dir
}
