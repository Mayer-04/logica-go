package main

import "fmt"

/*
Escribe una función `argumentsLength` que devuelva el número de argumentos se le pasan.

Ejemplo:
- Input: args = [5] - Output: 1
- Input: args = [true, 10, "3"] - Output: 3
*/

func main() {
	result := argumentsLength("mayer", "lucas")
	fmt.Println(result)

	// Ejemplo 2
	// El operador ... (variadic) en Go solo puede desempaquetar slices, no arrays.
	arr := [1]int{5}
	fmt.Println(argumentsLength(arr[:]...)) // Convertimos el array a slice y luego desempaquetamos

	// Ejemplo 3
	arr2 := []any{true, 10, "3"}
	fmt.Println(argumentsLength(arr2...)) // Desempaquetar el slice
}

func argumentsLength[T any](args ...T) int {
	return len(args)
}
