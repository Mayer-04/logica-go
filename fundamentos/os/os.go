package main

import (
	"fmt"
	"os"
)

/*
* Paquete os
Es uno de los paquetes más importantes del lenguaje Go.
El paquete `os` proporciona acceso al sistema operativo (OS) y a la entrada/salida (I/O) del sistema.

- Manipulación de archivos y manipulación de variables de entorno.


os.Stdout
os.Stderr
os.Stdin
*/

func main() {
	// Obtener el directorio de trabajo actual.
	dir, err := os.Getwd()
	if err != nil {
		// Devuelve un error envuelto con un mensaje explicativo si no se puede obtener el directorio.
		return
	}
	// Devuelve la ruta del directorio actual si no hubo errores.
	fmt.Println("Directorio de trabajo actual:", dir)

	// Obtener una variable de entorno.
	// valueEnv := "PORT"
	value := os.Getenv("PORT")
	fmt.Println(value)

	// Establecer el valor de una variable de entorno.
	var newEnv string
	if err := os.Setenv(newEnv, "10"); err != nil {
		fmt.Println(err)
	}
	fmt.Println(newEnv)

	// Eliminar una variable de entorno.
	if err := os.Unsetenv(newEnv); err != nil {
		fmt.Println(err)
	}

	// Verificar si una variable de entorno existe.
	_, ok := os.LookupEnv(newEnv)
	if ok {
		fmt.Println("variable de entorno existe")
	} else {
		fmt.Println("la variable de entorno no existe")
	}
}
