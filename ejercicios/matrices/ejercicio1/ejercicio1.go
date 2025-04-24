package main

import (
	"fmt"
)

/*
Dada una matriz (3x3), calcula la diferencia `absoluta` entre las sumas de sus diagonales.

CONSIDERACIONES:
- La diferencia absoluta entre dos números es siempre positiva o cero.
- Se calcula como: |a - b|, es decir, si `a` es mayor o menor que `b`, el resultado siempre es positivo.
- Si el número es negativo, se convierte en positivo, y si ya es positivo o es cero, se queda igual.
- Multiplicar un número negativo por -1 siempre resulta en un número positivo.

Por ejemplo, la matriz cuadrada se muestra a continuación:

1 2 3
4 5 6
9 8 9

- La diagonal de izquierda a derecha = 1 + 5 + 9 = 15 .
- La diagonal de derecha a izquierda = 3 + 5 + 9 = 17.
- Su diferencia absoluta es = |15 - 17| = 2.

Entrada:
[][]int: Es la matriz de números enteros.
Salida:
Devolver la diferencia absoluta entre las sumas de sus diagonales.
*/

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	fmt.Println(diagonalDifference(arr))  // Output: 2
	fmt.Println(diagonalDifference2(arr)) // Output: 2
	fmt.Println(diagonalDifference3(arr)) // Output: 2
}

func diagonalDifference(arr [][]int) int {
	sumaIzq := 0
	sumaDer := 0

	for i := range arr {
		sumaIzq += arr[i][i]
		sumaDer += arr[i][len(arr)-1-i]
	}

	// absoluta := int(math.Abs(float64(sumaIzq - sumaDer)))
	if sumaIzq > sumaDer {
		return sumaIzq - sumaDer
	}
	return sumaDer - sumaIzq
}

func diagonalDifference2(arr [][]int) int {
	sumaIzq := 0
	sumaDer := 0

	for i, fila := range arr {
		sumaIzq += fila[i]
		sumaDer += fila[len(arr)-1-i]
	}

	sumaTotal := sumaIzq - sumaDer
	if sumaTotal < 0 {
		sumaTotal *= -1
	}
	return sumaTotal
}

func diagonalDifference3(arr [][]int) int {
	sumaIzq := 0
	sumaDer := 0

	for i, fila := range arr {
		sumaIzq += fila[i]
		sumaDer += fila[len(arr)-1-i]
	}

	sumaTotal := sumaIzq - sumaDer
	if sumaTotal < 0 {
		sumaTotal = -sumaTotal // -(-2) = 2
	}
	return sumaTotal
}
