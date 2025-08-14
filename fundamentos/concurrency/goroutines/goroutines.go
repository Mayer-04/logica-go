package main

import (
	"fmt"
	"time"
)

/*
* Goroutines en Go

Las goroutines son `hilos ligeros` administrados por el `runtime de Go` que permiten ejecutar funciones
de manera concurrente.

* Definición clara:
Una goroutine es una función que se ejecuta en un hilo independiente del hilo principal.
No es un thread del sistema operativo, sino un hilo ligero gestionado por el runtime de Go.

* Características principales:
- Eficiencia de memoria: Comienzan con solo 2KB de pila, permitiendo ejecutar miles de tareas
  concurrentes sin sobrecargar el sistema.
- Sencillez de uso: Para iniciar una goroutine, solo se añade la palabra clave 'go' antes de llamar a la función.
- Función principal: El programa comienza con 'main' como goroutine principal.
- Comportamiento independiente: La goroutine 'main' no espera automáticamente a que otras goroutines terminen.
- Sin garantía de orden: Las goroutines pueden ejecutarse en cualquier orden.
- Sincronización recomendada: Es mejor usar canales (channels) o el paquete 'sync' en lugar de 'time.Sleep'.

* Scheduler (Planificador) del runtime de Go
Cuando lanzas una goroutine, el runtime de Go determina cuándo y cómo se ejecutará
utilizando el modelo GMP (Goroutine, Machine, Processor).

* Funcionamiento del scheduler de Go
El planificador administra la ejecución de goroutines mediante colas y un modelo basado en tres componentes:

- G (Goroutine): Representa la tarea a ejecutar.
- M (Machine): Son los hilos del sistema operativo que ejecutan las goroutines.
- P (Processor): Son entidades lógicas que gestionan las colas de goroutines y las asignan
a los hilos del sistema operativo.

* Proceso de planificación:
- Cada Processor (P) mantiene su propia cola de goroutines.
- Los Processors se asocian con Machines (hilos del SO) para ejecutar las goroutines.
- Múltiples goroutines se ejecutan sobre un número menor de hilos del sistema operativo,
gracias a la gestión eficiente del scheduler.
- Si una goroutine se bloquea, el scheduler puede reasignar el hilo (M) a otro Processor
y ejecutar otra goroutine.
- El runtime redistribuye goroutines entre colas locales y globales para balancear
la carga de trabajo eficientemente.

✅ IMPORTANTE: Si la goroutine principal ('main') termina antes que una goroutine secundaria,
esta última también terminará inmediatamente, incluso si no ha completado su tarea.
*/

func main() {
	// Iniciamos dos goroutines que se ejecutarán concurrentemente.
	// El orden de ejecución de estas goroutines no es garantizado.
	go goroutine1()
	go goroutine2()

	// Este mensaje se imprime desde la goroutine principal `main`.
	fmt.Println("Programa principal")

	// Esperamos 2 segundos para dar tiempo a que las goroutines completen su ejecución.
	// Nota: Este método es simple pero no ideal para sincronizar goroutines.
	time.Sleep(2 * time.Second)

	// Esta tercera goroutine no tendrá tiempo para ejecutarse porque 'main' terminará inmediatamente después.
	go goroutine3()
}

func goroutine1() {
	fmt.Println("Soy la primera goroutine")
}

func goroutine2() {
	fmt.Println("Soy la segunda goroutine")
}

func goroutine3() {
	fmt.Println("Soy la tercera goroutine, pero no me verás porque el programa principal termina antes")
}
