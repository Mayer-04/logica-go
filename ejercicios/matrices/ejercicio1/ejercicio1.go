package main

import "fmt"

/*
Dada una matriz cuadrada (3x3), calcula la diferencia `absoluta` entre las sumas de sus diagonales.

Por ejemplo, la matriz cuadrada se muestra a continuación:

1 2 3
4 5 6
9 8 9

- La diagonal de izquierda a derecha = 1 + 5 + 9 = 15 .
- La diagonal de derecha a izquierda = 3 + 5 + 9 = 17.
- Su diferencia absoluta es = |15 - 17| = 2.

[][]int: una matriz de números enteros
Devolvera la diferencia absoluta entre las sumas de sus diagonales
*/

func main() {
	arr := [][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}

	fmt.Println(diagonalDifference(arr))
}

func diagonalDifference(arr [][]int) int {
	var suma1 int
	var suma2 int
	var result int
	for i := 0; i < len(arr); i++ {
		suma1 += arr[i][i]
		suma2 += arr[i][len(arr)-1-i]
		if suma1 > suma2 {
			result = suma1 - suma2
		} else {
			result = suma2 - suma1
		}
	}
	return result
}
