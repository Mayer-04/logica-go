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

En lugar de usar técnicas tradicionales de sincronización (como mutex) para proteger memoria compartida,
los canales permiten que las gorrutinas compartan memoria a través de la comunicación.

* Tipos de canales
- chan T: Canal bidireccional (puede enviar y recibir valores de tipo T).
- chan<- T: Canal de solo envío (solo puede enviar valores).
- <-chan T: Canal de solo recepción (solo puede recibir valores).

- Los canales se crean con la función `make()` y su tipo de dato es `chan`.
- El operador `<-` se usa para enviar datos a un canal (canal <- dato) y también
para recibir datos de un canal (dato <- canal).
- Un canal declarado pero no inicializado es `nil`.
- Enviar como recibir desde un canal `nil` bloqueara el programa.
- Un canal se bloquea si no hay una goroutine disponible para recibir o enviar datos.
- La función `close()` se usa para cerrar un canal, indicando que no se enviarán más datos a través de él.
- Enviar datos a un canal que se ha cerrado con `close()` causará un `panic`.
- Recibir datos de un canal cerrado devolvera el `valor cero` del tipo del canal.
- NO es obligatorio cerrar un canal, el recolector de basura se encargara de esto cuando no se utilice el canal.

* Tipos de canales: con y sin búfer
En Go, los canales pueden ser con búfer o sin búfer, y su comportamiento varía según el tipo:

- Canal sin búfer: La operación de envío se bloquea hasta que otro goroutine reciba el valor.
Es decir, el envío y la recepción deben ocurrir al mismo tiempo.
- Canal con búfer: Puede almacenar varios valores sin necesidad de que otro goroutine los reciba de inmediato.
Solo se bloquea cuando el búfer está lleno.

* IMPORTANTE:
- Evita el deadlock: Un deadlock ocurre cuando todas las goroutines están esperando unas a otras
y ninguna puede continuar. Este es un error común al trabajar con canales mal sincronizados.
- Canales unidireccionales: Son canales que se pueden usar exclusivamente para enviar o para recibir datos.
Son útiles para restringir el acceso a un canal

* Señalizar
Significa notificar a una goroutine (o a varias) que un evento ha ocurrido, sin necesidad de transferir datos.

* Notificar o señalizar eventos:
- Cuando usamos canales (chan) en Go, a menudo queremos comunicar información entre goroutines.
Sin embargo, en algunos casos no necesitamos enviar datos, sino solo indicar que algo ha ocurrido.
- 'chan struct{}' nos permite señalizar eventos sin desperdiciar memoria.
- Es más eficiente que 'chan bool' porque no consume memoria.

* Cerrar un canal: close()
- Solo se pueden cerrar canales que NO sean de solo recepción.
- Cerrar un canal nil o ya cerrado causa panic.


* Transferencia de Valores
Copia de Valores

- Los valores se copian cuando se transfieren entre gorrutinas.
- Para canales con buffer: dos copias (emisor→buffer, buffer→receptor).
- Para canales sin buffer: una copia (emisor→receptor directamente).
- Solo se copia la parte directa del valor.
*/

func main() {
	// Valor cero de un canal es nil.
	var channel chan int             // ch3 == nil
	fmt.Println("channel:", channel) // Output: <nil>

	// Creamos un canal de tipo string sin búfer.
	// Esto significa que las operaciones de envío se bloquearán hasta que alguien reciba el valor.
	ch := make(chan string)

	// * Enviamos datos al canal desde una goroutine.
	go func() {
		time.Sleep(1 * time.Second) // Simulamos una operación que tarda 1 segundo en completarse.
		ch <- "Hello, World!"       // Enviamos el mensaje "Hello, World!" al canal `ch`.
	}()

	// * Recibimos un valor del canal.
	// `msg` almacena el valor recibido desde el canal.
	// `ok` indica si la recepción fue exitosa (true) o si el canal fue cerrado (false).
	msg, ok := <-ch

	// Imprimimos el valor recibido y el estado del canal.
	fmt.Println("mensaje:", msg)      // Output: Hello, World!
	fmt.Println("canal abierto:", ok) // Output: true

	// * Canales unidireccionales.
	num := make(chan int)

	// Este canal solo puede recibir datos.
	go receive(num)

	// Este canal solo puede enviar datos.
	send(num)

	//* Creando un canal con búfer.
	// Lo que define el tamaño del búfer es cuántos valores pueden estar en el canal al mismo tiempo
	// sin que el envío se bloquee.
	ch2 := make(chan int, 2) // Puede almacenar hasta 2 enteros al mismo tiempo antes de bloquearse.

	// Enviando datos al canal con búfer.
	// Estas operaciones no se bloquean porque el canal aún tiene espacio disponible.
	// El envío a un canal con búfer solo se bloquea cuando el búfer está lleno.
	ch2 <- 4
	ch2 <- 2

	// ch2 <- 6 // Esta línea causaría un bloqueo si se descomenta,
	// porque el canal ya ha alcanzado su capacidad máxima (2 valores).

	// Recibiendo datos del canal.
	// Las operaciones de recepción solo se bloquearán cuando el búfer esté vacío.
	fmt.Println(<-ch2) // Imprime el primer valor almacenado en el canal (4).
	fmt.Println(<-ch2) // Imprime el segundo valor almacenado en el canal (2).

	// * Ejemplo de uso de un canal con búfer y su posterior cierre.
	// Creamos un canal de enteros con capacidad para almacenar hasta 2 valores sin bloquearse.
	ch3 := make(chan int, 2)

	// Iniciamos una goroutine que enviará datos al canal.
	go func() {
		for i := 1; i <= 3; i++ {
			ch3 <- i // Enviamos valores al canal. El envío se bloquea solo si el búfer está lleno.
		}

		// Cerramos el canal una vez que se han enviado todos los valores.
		// Esto es importante para que el receptor sepa que no habrá más datos.
		close(ch3)
	}()

	// Iteramos sobre todos los valores recibidos del canal `ch3`.
	// El bucle continúa leyendo hasta que el canal sea cerrado.
	// Cada vez que se lee un valor del canal, se libera un espacio en el búfer,
	// lo que permite que otros envíos (bloqueados o futuros) puedan completarse.
	for value := range ch3 {
		// Leemos un valor del canal y lo imprimimos.
		fmt.Println("valor:", value) // Output: 1, 2, 3
	}

	//* Canales de notificación (o señalización).
	// Se utilizan para informar a otras goroutines que una tarea ha finalizado o que un evento ha ocurrido.
	// A diferencia de otros canales, no se usan para enviar datos, sino únicamente como mecanismo de sincronización.

	// Usamos `chan struct{}` en lugar de `chan bool` porque `struct{}` no ocupa memoria adicional.
	// Es una convención idiomática en Go para este tipo de casos.
	done := make(chan struct{})

	go func() {
		fmt.Println("Goroutine en ejecución")
		// Cuando `close` es ejecutado, `main()` detecta que el canal se cerró y puede seguir ejecutando.
		close(done) // 🔴 No enviamos datos, solo cerramos el canal.
	}()

	// No recibe datos, solo detecta que el canal fue cerrado.
	<-done // ⏳ Se desbloquea cuando el canal se cierra
	fmt.Println("Finalizado")
}

// send envía un valor entero a través del canal.
// El parámetro `num` es un canal unidireccional de solo envío (chan<- int),
// lo que significa que esta función solo puede enviar datos a ese canal,
// pero no puede recibir desde él.
func send(num chan<- int) {
	num <- 10
}

// receive recibe un valor entero desde el canal y lo imprime.
// El parámetro `num` es un canal unidireccional de solo recepción (<-chan int),
// por lo tanto, esta función solo puede leer datos del canal, pero no puede enviar a él.
func receive(num <-chan int) {
	fmt.Println(<-num)
}
