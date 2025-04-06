package main

import (
	"fmt"
	"slices"
)

/*
* Dividir un slice:
Escribe una función que tome un slice de enteros y lo divida en dos slices,
uno que contenga los primeros n elementos y otro que contenga el resto.
La función debería devolver ambos slices.
*/

func main() {
	enteros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	primeros, resto := dosSlices(enteros, 4)

	fmt.Println("primeros:", primeros)
	fmt.Println("resto:", resto)
}

func dosSlices(enteros []int, n int) ([]int, []int) {
	// Si n es mayor que la longitud del slice, devolver el slice completo y un slice vacío.
	if n > len(enteros) {
		return enteros, []int{}
	}

	// Slicing del slice de enteros para obtener los primeros n elementos.
	primeros := enteros[:n]

	var resto []int

	// Recorremos el slice de enteros
	for _, value := range enteros {
		// método `Contains` del paquete slices que verifica si en "primeros" estan los valores de enteros.
		result := slices.Contains(primeros, value)

		// Si en "primeros" no estan los valores del slice de enteros entonces los agrego a la variable resto.
		if !result {
			resto = append(resto, value)
		}

	}

	// Retorno los dos slices, primeros y resto
	return primeros, resto
}

func DividirSlice(slice []int, n int) ([]int, []int) {
	// Verificar si n es mayor que la longitud del slice
	if n > len(slice) {
		return slice, []int{}
	}

	// Slicing del slice de enteros para obtener los primeros n elementos.
	primeros := slice[:n]
	// Slicing del slice de enteros para obtener el resto de los elementos.
	resto := slice[n:]

	return primeros, resto
}
