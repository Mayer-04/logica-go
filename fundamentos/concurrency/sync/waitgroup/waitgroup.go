package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

/*
* Sincronización de Goroutines con WaitGroup

El paquete "sync" proporciona primitivas para sincronizar goroutines, entre ellas el `WaitGroup`.

El objetivo de usar un `WaitGroup` es asegurarse de que todas las goroutines terminen
antes de que el programa principal continúe, pero no se controla el orden en que se ejecutan las goroutines.

- `sync.WaitGroup`: Estructura que permite esperar a que un grupo de goroutines termine su ejecución.
- `Add(delta int)`: Incrementa el contador del WaitGroup por el valor de `delta`.
- `Done()`: Decrementa el contador del WaitGroup en 1. Debe ser llamado cuando una goroutine termina.
- `Wait()`: Bloquea la ejecución hasta que el contador del WaitGroup llega a 0, es decir,
hasta que todas las goroutines hayan llamado a `Done()`.
*/

func main() {
	waitGroup()
	fetchAllAPIs()
}

// waitGroup inicializa un WaitGroup, lanza múltiples goroutines y espera a que todas completen su ejecución.
func waitGroup() {
	// Crea una instancia de WaitGroup para manejar la sincronización de las gorutinas.
	var wg sync.WaitGroup

	// Definimos una constante para indicar el número de gorutinas que vamos a lanzar.
	const numGoroutines = 5
	// Incrementa el contador del WaitGroup en 5, indicando que vamos a esperar a que 5 gorutinas terminen.
	// NOTA: Llama al método `Add()` antes de que se llame al método `Done()`, no utilizarlo dentro de las gorutinas.
	wg.Add(numGoroutines)

	// Lanza 5 gorutinas que imprimen su índice antes de finalizar.
	for i := 1; i <= numGoroutines; i++ {
		go func(i int) {
			// Decrementa el contador del WaitGroup cuando la goroutine finaliza.
			// Cuando usas defer, te aseguras de que wg.Done() se ejecute justo antes de que la goroutine termine.
			// Nos ayuda a no preocuparnos en donde colocar el wg.Done() en el código.
			defer wg.Done()

			// Imprime el índice de la goroutine actual.
			fmt.Println("Goroutine:", i)
		}(i)
	}

	// Bloquea el hilo principal (main) hasta que el contador del WaitGroup sea 0,
	// es decir, hasta que todas las goroutines que hemos añadido con `Add()` hayan terminado.
	wg.Wait()

	// Imprime un mensaje indicando que todas las goroutines han finalizado.
	fmt.Println("Todas las goroutines se han completado!")
}

//* Consumiendo múltiples APIs con WaitGroup:

// urls de las 3 APIs que queremos solicitar.
var urls = [3]string{
	"https://jsonplaceholder.typicode.com/posts/1",
	"https://rickandmortyapi.com/api/character",
	"https://jsonplaceholder.typicode.com/users",
}

func fetchAllAPIs() {
	// Inicializa un WaitGroup para esperar hasta que todas las solicitudes a las APIs hayan finalizado.
	var wg sync.WaitGroup

	// Incrementa el contador del WaitGroup con la longitud del slice de APIs.
	wg.Add(len(urls))

	// Lanza una gorutina por cada API en el slice.
	for _, url := range urls {
		go func(url string) {
			// Decrementa el contador del WaitGroup cuando cada gorutina finaliza.
			defer wg.Done()
			// Llama a la función requestAPI para realizar la solicitud HTTP.
			data, err := RequestAPI(url)
			if err != nil {
				fmt.Println(err)
				return
			}
			// Imprime los datos recibidos de la APIs como una cadena de texto.
			fmt.Println(string(data))
		}(url) // Pasamos la URL actual de la API a la función anonima (gorutina).
	}

	// Espera a que todas las gorutinas hayan completado su ejecución antes de continuar.
	wg.Wait()
}

// RequestAPI realiza una solicitud HTTP GET a una URL proporcionada y devuelve el cuerpo de la respuesta
// como un slice de bytes o un error detallado.
//
// El cuerpo de la respuesta se devuelve como un slice de bytes.
// Si se produce un error, se devuelve un error con una descripción detallada.
func RequestAPI(url string) ([]byte, error) {
	// Verifica que la URL no esta vacía.
	if url == "" {
		return nil, fmt.Errorf("la URL proporcionada no puede estar vacía")
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fallo al realizar la solicitud a %q: %w", url, err)
	}

	// Asegura que el cuerpo de la respuesta HTTP se cierre al finalizar la función.
	// Este patrón asegura que el cuerpo se cierre, pero no "drena" el contenido del cuerpo.
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("no se pudo cerrar el cuerpo de la respuesta: %v\n", err)
		}
	}()

	// defer func() {
	// 	// Drenar: Leer el cuerpo por completo, incluso si no necesitamos su contenido.
	// 	// `io.Copy` es una función que copia los datos desde un origen (src) a un destino (dst).
	// 	// `io.Discard` es el destino, que descarta cualquier dato recibido. Es una especie de "basurero".
	// 	// `resp.Body` es el origen, que contiene los datos del cuerpo de la respuesta HTTP.
	// 	_, err := io.Copy(io.Discard, resp.Body)
	// 	if err != nil {
	// 		log.Printf("error al copiar el cuerpo de la respuesta: %v\n", err)
	// 	}

	// 	// Cerramos el cuerpo de la respuesta para liberar los recursos asociados con la conexión.
	// 	err = resp.Body.Close()

	// 	// Si ocurre un error al cerrar el cuerpo, lo registramos.
	// 	if err != nil {
	// 		log.Printf("no se pudo cerrar el cuerpo de la respuesta: %v\n", err)
	// 	}
	// }()

	// Verifica que el estado de la respuesta HTTP sea 200 OK.
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("el código de estado de respuesta es %d, esperado 200", resp.StatusCode)
	}

	// Lee el cuerpo de la respuesta HTTP y devuelve los datos como un slice de bytes.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %w", err)
	}

	// Devuelve los datos leídos como un slice de bytes y un error nil, indicando que la operación fue exitosa.
	return body, nil
}
