package main

import "fmt"

/*
* Eliminar elementos de un slice:
Escribe una función que tome un slice de enteros y elimine todos los números impares del slice.
Asegúrate de que el slice resultante mantenga su longitud y capacidad original.
*/

func main() {
	enteros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("enteros: %d, capacidad: %d\n", len(enteros), cap(enteros)) // enteros: 9, capacidad: 9

	result := eliminarElementos(enteros)

	fmt.Printf("pares: %d, capacidad: %d\n", len(result), cap(result)) // pares: 4, capacidad: 9
	fmt.Println("Resultado:", result)                                  // Resultado: [2 4 6 8]
}

func eliminarElementos(enteros []int) []int {
	var pares = make([]int, 0, len(enteros))

	for _, value := range enteros {
		if value%2 == 0 {
			pares = append(pares, value)
		}
	}

	return pares
}
