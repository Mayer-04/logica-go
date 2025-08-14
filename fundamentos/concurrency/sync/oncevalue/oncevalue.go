package main

import (
	"fmt"
	"sync"
)

/*
* sync.OnceValue

sync.OnceValue es una función en Go, introducida en la versión 1.21, que garantiza que una función
se ejecute exactamente una vez y reutiliza el valor retornado en llamadas subsecuentes.

- No importa cuántas veces se llame a la función devuelta por OnceValue, la función original solo se ejecutará una vez.
- Es más conveniente cuando la función que estás ejecutando retorna un valor que necesitas reutilizar,
ya que encapsula la lógica de ejecución única y la devolución del valor.
- Previene 'condiciones de carrera' al inicializar recursos compartidos entre varias goroutines,
asegurando que múltiples goroutines que llamen a la función simultáneamente obtengan el mismo valor sin reinicializar.
- Después de que la función asociada se haya ejecutado una vez, no puede ser "reseteada" o ejecutada nuevamente.
Cualquier llamada futura simplemente reutiliza el valor retornado previamente.

* Método `OnceValue`:
- Devuelve una función que, cuando es llamada, ejecuta la función f (en este caso `LoadMessage`) una sola vez.
Las subsecuentes invocaciones a la función retornarán el valor de la primera ejecución de f.
*/

// LoadMessage es una función que imprime un mensaje y retorna un valor.
// En este caso, solo se imprimirá "Loading message..." la primera vez que se llame GetMessage().
func LoadMessage() string {
	fmt.Println("Loading message...")
	return "Hello, Go!"
}

// LoadConfig carga una configuración desde un archivo (simulación).
func LoadConfig() map[string]string {
	fmt.Println("Cargando configuración...")
	return map[string]string{"ambiente": "producción"}
}

// GetMessage garantiza que LoadMessage se ejecute solo una vez y retorna el mensaje cargado.
var GetMessage = sync.OnceValue(LoadMessage)

// config es una variable que almacena la configuración cargada.
var config = sync.OnceValue(LoadConfig)

// La función main llama a GetMessage varias veces, pero LoadMessage() solo se ejecuta la primera vez.
// Las llamadas subsecuentes a GetMessage() simplemente reutilizan el valor retornado la primera vez.
func main() {
	// Primera llamada: se ejecuta LoadMessage y se imprime "Loading message...".
	message := GetMessage()
	fmt.Println(message)

	// Llamadas subsecuentes: reutilizan el valor retornado previamente.
	fmt.Println(GetMessage()) // No ejecuta LoadMessage, solo retorna "Hello, Go!"
	fmt.Println(message)      // No ejecuta LoadMessage, solo retorna "Hello, Go!"

	// Primera vez que se accede a la configuración:
	cfg := config()
	fmt.Println(cfg["ambiente"])

	// Segunda vez: se reutiliza el valor cargado previamente.
	cfg = config()
	fmt.Println(cfg["ambiente"])
}
