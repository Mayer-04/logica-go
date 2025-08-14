package main

import "fmt"

/*
Crea una matriz de cualquier tamaño.
Llena las columnas impares con números que comienzan en 2 y aumentan de dos en dos (2, 4, 6, etc.).
Llena las columnas pares con el número 0 en todas sus filas.

Ejemplo:

matrix := [][]int{
	{2, 0, 4},
	{0, 6, 0},
	{8, 0, 10},
}
*/

func main() {
	matrix := [][]int{
		{2, 0, 4},
		{0, 6, 0},
		{8, 0, 10},
	}

	num := 2
	for i := range matrix {
		filas := matrix[i]
		for j := range filas {
			if j%2 == 0 {
				matrix[i][j] = num
				num += 2
			}
		}
	}

	fmt.Println(matrix) // Output: [[2 0 4] [6 6 8] [10 0 12]]
}
