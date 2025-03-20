package main

import (
	"fmt"
	"time"
)

/*
* Channels: Canales en Go

📖 FRASE: No se comunique compartiendo memoria, comparta memoria comunicándose.

Los canales en Go permiten la comunicación y sincronización entre goroutines,
facilitando el intercambio seguro de datos.

- Los canales se crean con la función `make()`, y su tipo de dato es `chan`.
- El operador `<-` se usa para enviar datos a un canal (canal <- dato) y también
para recibir datos de un canal (dato <- canal).
- Un canal declarado pero no inicializado es `nil`.
- Enviar como recibir desde un canal `nil` bloqueara el programa.
- Un canal se bloquea si no hay una goroutine disponible para recibir o enviar datos.
- La función `close()` se usa para cerrar un canal, indicando que no se enviarán más datos a través de él.
- Enviar datos a un canal que se ha cerrado con `close()` causará un `panic`.
- Recibir datos de un canal cerrado devolvera el `valor cero` del tipo del canal.

* Existen canales con y sin búfer:
- Sin búfer: Los envíos y recepciones se bloquean hasta que la otra parte esté lista.
- Con búfer: Se pueden almacenar varios datos en el canal antes de que se bloquee.

* IMPORTANTE:
- Evitar el `deadlock`, que ocurre cuando todas las goroutines están bloqueadas, esperando unas de otras.
- Los canales unidireccionales se usan para limitar el uso de un canal a solo enviar o solo recibir datos,
lo cual puede ayudar a evitar errores.

* Señalizar
Significa notificar a una goroutine (o a varias) que un evento ha ocurrido, sin necesidad de transferir datos.

* Notificar o señalizar eventos:
- Cuando usamos canales (chan) en Go, a menudo queremos comunicar información entre goroutines.
Sin embargo, en algunos casos no necesitamos enviar datos, sino solo indicar que algo ha ocurrido.
- 'chan struct{}' nos permite señalizar eventos sin desperdiciar memoria.
- Es más eficiente que 'chan bool' porque no consume memoria.
*/

func main() {
	// Creando un canal de tipo string.
	ch := make(chan string)

	//* Enviar datos al canal desde una goroutine.
	go func() {
		time.Sleep(1 * time.Second) // Simulamos una operación que 1 segundo en ejecutarse.
		ch <- "Hello, World!"       // Envía el mensaje "Hello, World!" al canal `ch`.
	}()

	//* Operación de recepción desde un canal.
	// `msg` es la variable donde se almacená el valor recibido desde el canal.
	// `ok` es un valor booleano que indica si la operación de recepción fue exitosa.
	msg, ok := <-ch
	fmt.Println("mensaje:", msg)      // Output: Hello, World!
	fmt.Println("canal abierto:", ok) // Output: true

	// * Canales unidireccionales.
	num := make(chan int)

	// Este canal solo puede recibir datos.
	go receive(num)

	// Este canal solo puede enviar datos.
	send(num)

	//* Creando un canal con búfer.
	// Este canal puede almacenar hasta 2 enteros antes de bloquearse.
	ch2 := make(chan int, 2)

	// Enviando datos al canal. No se bloquea porque tiene capacidad para almacenar dos valores.
	// Las operaciones de envío al canal solo se bloquearán cuando el búfer esté lleno.
	ch2 <- 4
	ch2 <- 2
	// ch2 <- 6 // Se bloquearía si se descomenta, ya que el búfer está lleno.

	// Recibiendo datos del canal.
	// Las operaciones de recepción solo se bloquearán cuando el búfer esté vacío.
	fmt.Println(<-ch2) // Imprime el primer valor almacenado en el canal (4).
	fmt.Println(<-ch2) // Imprime el segundo valor almacenado en el canal (2).

	// * Ejemplo de uso de canal cerrado.
	// Creando un canal de enteros.
	ch3 := make(chan int, 2)

	go func() {
		for i := 1; i <= 3; i++ {
			ch3 <- i // Enviando valores al canal.
		}
		close(ch3) // Cerrando el canal después de enviar todos los valores.
	}()

	// Iterando todos los valores pasados por el canal `ch3`.
	// Recibimos los datos del canal hasta que esté cerrado.
	for value := range ch3 {
		// Leemos cada uno de los valores del canal y los imprimimos.
		fmt.Println("valor:", value) // Output: 1, 2, 3
	}

	//* Canales de notificación.
	// Se utilizan cuando se necesita notificar o señalizar a otras goroutines de que algo ha ocurrido.
	// No se usan para enviar datos entre goroutines.

	// Se usa 'chan struct{}' en lugar de 'chan bool' para reducir el uso de memoria.
	done := make(chan struct{})

	go func() {
		fmt.Println("Goroutine en ejecución")
		// Cuando `close` es ejecutado, `main()` detecta que el canal se cerró y puede seguir ejecutando.
		close(done) // 🔴 No enviamos datos, solo cerramos el canal
	}()

	// No recibe datos, solo detecta que el canal fue cerrado.
	<-done // ⏳ Se desbloquea cuando el canal se cierra
	fmt.Println("Finalizado")
}

// `send` envía un valor entero al canal.
// Esta función solo acepta un canal que envie datos.
func send(num chan<- int) {
	num <- 10
}

// `receive` recibe un valor entero del canal y lo imprime.
// Esta función solo acepta un canal que reciba datos.
func receive(num <-chan int) {
	fmt.Println(<-num)
}
