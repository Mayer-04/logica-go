package main

import "fmt"

/*
* Elimina duplicados:
Escribe un programa que reciba un slice de enteros y devuelva un nuevo slice sin números duplicados.
Por ejemplo, si la entrada es [1, 2, 2, 3, 4, 4, 5], la salida debería ser [1, 2, 3, 4, 5].
*/

func main() {
	num := []int{1, 2, 2, 3, 4, 4, 5}
	result := eliminarDuplicados(num)
	fmt.Println(result)
}

func eliminarDuplicados(nums []int) []int {
	// Crear un mapa para almacenar los números únicos
	numerosUnicos := make(map[int]bool)

	// Crear un slice para los números únicos
	resultado := make([]int, 0, len(numerosUnicos))

	// num := []int{1, 2, 2, 3, 4, 4, 5}
	// Recorrer el slice original
	for _, num := range nums {
		// Si el número no ha sido agregado aún al mapa, lo agregamos al resultado
		if !numerosUnicos[num] {
			resultado = append(resultado, num)
			numerosUnicos[num] = true
		}
	}

	// Retornar el slice sin duplicados
	return resultado
}
