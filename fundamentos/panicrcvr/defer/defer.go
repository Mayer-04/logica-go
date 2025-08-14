package main

import (
	"errors"
	"fmt"
	"os"
)

/*
* Defer
La palabra clave `defer` en Go se usa para retrasar la ejecución de una función
hasta que la función que la contiene haya terminado.

- Es como programar una tarea que se ejecuta justo antes de que la función finalice,
sin importar si termina normalmente o por un pánico.
- Es útil para tareas de limpieza: cerrar archivos, liberar recursos, cerrar conexiones, etc.
*/

func main() {
	// En este ejemplo, se retrasa la impresión de "hello" hasta que termine la función main,
	// lo que significa que "hello" se imprimirá después de "world".
	defer fmt.Println("hello")
	fmt.Println("world")

	//* SALIDA:
	// world
	// hello

	doSomething()
	deferMultiple()
	deferWithRecover()

	// Los argumentos pasados a una función defer se evalúan en el momento en que se declara el defer,
	// no cuando se ejecuta. Aquí, el valor de `a` es capturado como 10.
	a := 10
	defer pushAnalytic(a) /// Se imprime 10, no 20.
	a = 20

	// Otro ejemplo del mismo comportamiento anterior:
	deferEvalExample()
}

// doSomething realiza una operación de lectura en un archivo y asegura que se cierre correctamente
// utilizando defer, incluso si ocurre un error.
func doSomething() (err error) {
	// Abre el archivo "prueba.txt" para lectura
	file, err := os.Open("prueba.txt")
	if err != nil {
		return err
	}

	// Utiliza defer para asegurarse de que el archivo se cierre al salir de la función.
	defer func() {
		// `errors.Join()` permite combinar múltiples errores en uno solo.
		// Si `file.Close()` retorna un error al intentar cerrar el archivo,
		// este error se combina con cualquier error anterior que haya ocurrido.
		err = errors.Join(err, file.Close())
	}()

	return nil
}

// DeferMultiple demuestra el uso de múltiples declaraciones `defer`.
// En Go, cuando se utilizan varios defer en una función,
// se ejecutan en orden LIFO (último en entrar, primero en salir),
// lo que significa que la última función diferida se ejecutará primero.
func deferMultiple() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	// SALIDA:
	// 3
	// 2
	// 1
}

// DeferWithRecover demuestra cómo manejar un `pánico` con `defer` y `recover`.
// `recover` es una función incorporada que permite recuperar la ejecución normal de un pánico.
func deferWithRecover() {
	defer func() {
		// `recover` se utiliza para recuperar la ejecución normal de un pánico.
		if r := recover(); r != nil {
			fmt.Println("Recuperado del pánico:", r)
		}
	}()

	fmt.Println("Este mensaje se imprimirá")
	panic("¡Ocurrió un pánico!")
	// Este mensaje no se imprimirá porque el pánico detiene la ejecución normal hasta que recover lo maneja.
	fmt.Println("Este mensaje no se imprimirá")
}

// pushAnalytic imprime el valor recibido.
// Sirve para demostrar que los argumentos en defer se evalúan inmediatamente.
func pushAnalytic(a int) {
	fmt.Println(a)
}

// deferEvalExample muestra cómo el valor capturado por defer no cambia
// aunque la variable original sí lo haga antes de que el defer se ejecute.
func deferEvalExample() {
	val := 5
	defer fmt.Println("Valor capturado en defer:", val)
	val = 10
	fmt.Println("Valor actual:", val)

	// SALIDA:
	// Valor actual: 10
	// Valor capturado en defer: 5
}
