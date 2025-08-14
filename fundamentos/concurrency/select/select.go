package main

import (
	"fmt"
	"time"
)

/*
* Select en Go

Se utiliza para escuchar múltiples canales al mismo tiempo.

- Es similar a un `switch` pero para canales.
- El primer canal con un valor listo para ser recibido se disparará y su cuerpo se ejecutará.
- Si hay varios canales listos al mismo tiempo, uno se elige al azar.
- El `default` se ejecuta inmediatamente si ningún canal tiene un valor listo (no espera).
- Usar `default` es útil cuando: No quieres que el programa se quede bloqueado esperando datos en un canal.

* Reglas del select
- No se permiten expresiones después de select.
- No se permite fallthrough.
- Cada `case` debe ser una operación de canal.
- Si hay operaciones no-bloqueantes, se elige una aleatoriamente.
- Si todas son bloqueantes y hay `default`, se ejecuta el default.
- Si no hay `default`, la gorrutina se bloquea.


* IMPORTANTE
- time.Tick(500 * time.Millisecond) // Envía valores cada 500ms (no usar en producción sin cerrar, puede causar fugas de memoria)
- time.After(2 * time.Second)       // Envía un único valor después de 2s
- time.Sleep(1 * time.Second)       // Bloquea la goroutine por 1s
*/

func main() {
	fmt.Println("Ejemplo 1: Selección básica entre canales")
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Ambas goroutines enviarán un mensaje al mismo tiempo.
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Mensaje desde la goroutine 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Mensaje desde la goroutine 2"
	}()

	// Como ambos canales estarán listos simultáneamente, la selección será aleatoria.
	select {
	case msg1, ok := <-ch1:
		fmt.Printf("Recibido de ch1: %s, canal abierto: %t\n", msg1, ok)
	case msg2, ok := <-ch2:
		fmt.Printf("Recibido de ch2: %s, canal abierto: %t\n", msg2, ok)
	}

	fmt.Println("\nEjemplo 2: Timeout con select")
	ch3 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second) // Simulamos una operación que tarda 3 segundos.
		ch3 <- "Operación completada"
	}()

	select {
	case msg := <-ch3:
		fmt.Println(msg) // Se ejecuta si recibe un mensaje del canal antes de 2 segundos.
		// Si en 2 segundos no ha llegado nada a ch3, cancela la espera y ejecuta este código.
	case <-time.After(2 * time.Second):
		// time.After devuelve un canal que envía un valor después del tiempo especificado.
		fmt.Println("Timeout: La operación tardó demasiado")
	}

	fmt.Println("\nEjemplo 3: Cláusula default (comportamiento no bloqueante)")
	ch4 := make(chan string)

	// Enviamos un mensaje después de un tiempo.
	go func() {
		time.Sleep(2 * time.Second)
		ch4 <- "Mensaje tardío desde ch4"
	}()

	// El default se ejecutará inmediatamente porque ch4 no está listo aún.
	select {
	case msg := <-ch4:
		fmt.Println(msg)
	default:
		fmt.Println("No hay mensajes disponibles en este momento (ejecución inmediata)")
	}

	fmt.Println("\nEjemplo 4: Uso típico de default en un bucle")
	ch5 := make(chan string)
	// El canal `done` se usara para indicar que el bucle debe terminar.
	done := make(chan struct{})

	// Goroutine que enviará un mensaje después de un tiempo.
	go func() {
		time.Sleep(2 * time.Second)
		ch5 <- "¡Mensaje finalmente disponible!"
		time.Sleep(1 * time.Second)
		done <- struct{}{}
	}()

	// Bucle que comprueba periódicamente si hay mensajes en el canal, sin bloquearse gracias al default.
loop:
	for {
		select {
		case msg := <-ch5:
			fmt.Println("Recibido:", msg)
		case <-done:
			fmt.Println("Terminando el bucle")
			break loop
		default:
			fmt.Println("Esperando... (haciendo otras tareas mientras tanto)")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
