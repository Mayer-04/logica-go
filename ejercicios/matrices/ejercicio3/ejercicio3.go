package main

import "fmt"

/*
* Suma de Filas y Columnas
Escribe un programa que reciba una matriz de enteros de 𝑛×𝑚 y calcule la suma de los elementos
de cada fila y de cada columna. El programa debe imprimir ambas sumas.

Entrada: Una matriz de enteros.
Salida: Las sumas de cada fila y de cada columna.

matriz := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

Salida esperada:
Suma de las filas: 6, 15, 24
Suma de las columnas: 12, 15, 18
*/

func main() {
	// Definimos una matriz de enteros
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Variables para almacenar las sumas de las filas y las columnas
	filas := len(matriz)       // 3
	columnas := len(matriz[0]) // 3

	sumaFilas := make([]int, filas)
	sumaColumnas := make([]int, columnas)

	// Calcular las sumas de las filas y las columnas
	for i := range filas {
		for j := range columnas {
			sumaFilas[i] += matriz[i][j]    // Sumar el valor a la fila correspondiente
			sumaColumnas[j] += matriz[i][j] // Sumar el valor a la columna correspondiente
		}
	}

	// Imprimir los resultados
	fmt.Println("Suma de las filas:", sumaFilas)
	fmt.Println("Suma de las columnas:", sumaColumnas)
}
