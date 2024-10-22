package main

import (
	"fmt"
	"time"
)

/*
* Goroutines
Las goroutines son hilos ligeros gestionados por el runtime de Go, que permiten ejecutar funciones de manera concurrente.
- Para iniciar una goroutine, solo necesitas agregar la palabra clave 'go' antes de la llamada a la función.
- La función `main` se ejecuta como una goroutine principal en todos los programas Go.
- Es importante recordar que la goroutine `main` no espera automáticamente a que otras goroutines finalicen.
- El orden de ejecución de las goroutines no está garantizado.
- Se recomienda usar canales (channels) o el paquete `sync` para sincronizar goroutines en vez de usar `time.Sleep`.

* Scheduler del runtime de Go - Planificador en tiempo de ejecución de Go
Cuando lanzas una goroutine, no se ejecuta inmediatamente en un hilo o CPU específico; en su lugar,
el runtime de Go maneja cuándo y cómo se ejecutará esa goroutine utilizando el modelo GMP (Goroutine, Machine, Processor).

* ¿Cómo funciona el scheduler del runtime de Go?
El runtime de Go administra la ejecución de las goroutines y decide cuál debe ejecutarse en un momento dado
utilizando un conjunto de colas y un modelo de programación basado en el concepto de Goroutines (G),
Máquinas (M), y Procesadores (P).

G: Representa las goroutines.
M: Son los hilos del sistema operativo que el runtime de Go usa para ejecutar las goroutines.
P: Son los procesadores (entidades lógicas) que gestionan las colas de goroutines y las asignan
a los hilos del sistema operativo (M) para su ejecución.

- Las goroutines son colocadas en colas de trabajo (work queues) asociadas a los P (Procesadores).
- Cada (P) tiene su propia cola de goroutines y se asocia con un (M) (hilo del sistema operativo) para ejecutar las goroutines.
- Múltiples goroutines se ejecutan sobre un número menor de hilos del sistema operativo (M),
gracias a la gestión eficiente del scheduler.
- Si una goroutine realiza una operación bloqueante,
el scheduler puede reasignar el (M) asociado a otro P y ejecutar otra goroutine en su lugar.
- El runtime también redistribuye goroutines entre colas locales y globales
para optimizar el rendimiento y balancear la carga de trabajo de manera eficiente.


NOTA: Si el programa principal (la goroutine `main`) termina antes de que una goroutine hija termine,
la goroutine hija también terminará inmediatamente, incluso si no ha completado su trabajo.
*/

func main() {
	// Iniciamos dos goroutines concurrentes.
	go goroutine1() // Inicia la primera goroutine
	go goroutine2() // Inicia la segunda goroutine

	// Este mensaje se imprime desde la goroutine `main`.
	fmt.Println("Soy el programa principal")

	// Usamos `time.Sleep` para darle tiempo a las goroutines para completar su ejecución.
	// Este es un método simple pero no ideal para sincronizar goroutines.
	time.Sleep(2 * time.Second)

	// Esta tercera goroutine no tiene tiempo para ejecutarse porque la goroutine `main` terminará.
	go goroutine3()
}

func goroutine1() {
	fmt.Println("Soy la primera goroutine")
}

func goroutine2() {
	fmt.Println("Soy la segunda goroutine")
}

func goroutine3() {
	fmt.Println("Soy la tercera goroutine, pero no me verás porque el programa principal termina antes.")
}
