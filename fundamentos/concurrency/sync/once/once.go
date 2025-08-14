package main

import (
	"fmt"
	"sync"
)

/*
* Sync.Once

Sync.Once es una estructura en Go que garantiza que una acción se ejecute exactamente una vez.

- No importa cuántas goroutines llamen a la función asociada, la función se ejecutará solo una vez.
- Previene 'condiciones de carrera' al inicializar recursos compartidos entre varias goroutines.
- Después de que la función asociada haya sido ejecutada, no puede ser "reseteda" o ejecutada nuevamente.
- `sync.Once` es útil para inicializar recursos de manera segura y eficiente.

* Método `Do`:
- Toma una función sin argumentos ni valores de retorno como argumento.
- La primera vez que se llama a `Do`, ejecutará la función pasada.
- Las siguientes llamadas a `Do` no volverán a ejecutar la función, independientemente del número de goroutines.
- Ideal para recursos que solo se deben configurar una vez, como bases de datos o conexiones de red.
*/

var once sync.Once

// printMessage es una función simple que imprime un mensaje en la consola.
// Este será el recurso que se inicializa una sola vez.
func printMessage() {
	fmt.Println("Hola Mundo!")
}

// Ejemplo de uso de `sync.Once` para inicializar un recurso una sola vez.
// En este caso, se utiliza para imprimir un mensaje en la consola solo una vez,
// sin importar cuántas goroutines lo intenten hacer.
func main() {
	// Creación de un WaitGroup para sincronizar las goroutines.
	var wg sync.WaitGroup
	// Espera a que terminen 3 goroutines.
	wg.Add(3)

	// Lanzamiento de 3 goroutines que intentarán ejecutar la función `printMessage`.
	for i := 1; i <= 3; i++ {
		go func() {
			once.Do(printMessage) // Asegura que `printMessage` se ejecute solo una vez.
			wg.Done()             // Marca la goroutine como completa.
		}()
	}
	// Espera a que todas las goroutines terminen antes de finalizar el programa.
	wg.Wait()
}
