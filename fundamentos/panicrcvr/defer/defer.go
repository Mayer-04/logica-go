package main

import (
	"errors"
	"fmt"
	"os"
)

/*
* Defer
Es una palabra clave en Go que se utiliza para posponer la ejecución de una función,
hasta que la función que la rodea haya terminado.

- Es como programar una tarea para que se ejecute justo antes de que la función finalice,
sin importar cómo salga de la función (ya sea por un retorno normal o un pánico).
- Es útil para tareas de limpieza, como cerrar archivos, liberar recursos y cerrar una conexión a una base de datos.
*/

func main() {
	// En este ejemplo, se pospone la impresión de "hello" hasta que la función `main` termine,
	// lo que significa que se imprimirá después de "world".
	defer fmt.Println("hello")
	fmt.Println("world")

	// Salida:
	// world
	// hello

	doSomething()
	deferMultiple()
	deferWithRecover()

	// Los valores de las variables usadas en un defer se capturan en el momento en que se declara el defer.
	a := 10
	defer pushAnalytic(a) // Captura el valor de a en este momento. Se imprime `10` no 20.
	a = 20
}

// * DoSomething realiza una operación de lectura en un archivo y asegura que se cierre correctamente
// utilizando defer, incluso si ocurre un error.
func doSomething() (err error) {
	// Abre el archivo "prueba.txt" para lectura
	file, err := os.Open("prueba.txt")
	if err != nil {
		return err
	}

	// Utiliza defer para asegurarse de que el archivo se cierre al salir de la función.
	defer func() {
		// errors.Join permite combinar múltiples errores en uno solo.
		// Si file.Close retorna un error al intentar cerrar el archivo,
		// este error se combina con cualquier error anterior que haya ocurrido.
		err = errors.Join(err, file.Close())
	}()

	return nil
}

// DeferMultiple demuestra el uso de múltiples declaraciones defer.
// En Go, cuando se utilizan varios defer en una función,
// se ejecutan en orden LIFO (último en entrar, primero en salir),
// lo que significa que la última función diferida se ejecutará primero.
func deferMultiple() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	// Salida:
	// 3
	// 2
	// 1
}

// DeferWithRecover demuestra cómo manejar un `pánico` con `defer` y `recover`.
// `recover` es una función incorporada que permite recuperar la ejecución normal de un pánico.
func deferWithRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado del pánico:", r)
		}
	}()

	fmt.Println("Este mensaje se imprimirá")
	panic("¡Ocurrió un pánico!")
	fmt.Println("Este mensaje no se imprimirá")
}

// Evaluación inmediata de argumentos
func pushAnalytic(a int) {
	fmt.Println(a)
}
