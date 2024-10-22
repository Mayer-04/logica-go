package main

import "fmt"

//* recover(): Recuperar la ejecución normal de la aplicación por un "panic"
//!NOTE: Antes de ver el mensaje del panic se ve el mensaje de la función deferida "defer"
// Cuando se ejecuta panic, la ejecución del programa se detiene excepto las funciones de "defer"
// las funciones diferidas (defers) aún se ejecutarán antes de que el programa termine completamente

func main() {
	defer fmt.Println("Esto se imprimirá al final") // orden de ejecución: 2
	fmt.Println("Comienzo del programa")            // orden de ejecución: 1
	panic("¡Algo salió mal!")                       // orden de ejecución: 3
}
