package main

import "fmt"

/*
* Panic: Pánico

📖 FRASE: No entres en pánico.

- Cuando se ejecuta panic, la ejecución del programa se detiene excepto las funciones de "defer"
*/

func main() {
	defer fmt.Println("Esto se imprimirá al final") // orden de ejecución: 2
	fmt.Println("Comienzo del programa")            // orden de ejecución: 1
	panic("¡Algo salió mal!")                       // orden de ejecución: 3
}
