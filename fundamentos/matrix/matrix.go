package main

import "fmt"

/*
* Matrices en Go
En Go, las matrices pueden definirse tanto con `arreglos` como con `slices`.
La diferencia clave es que los arreglos tienen un tamaño fijo, mientras que los slices son dinámicos.
- Una matriz (array bidimensional) con arreglos se declara como `[m][n]Tipo`,
donde `m` es el número de filas y `n` el número de columnas. Esto crea una matriz de tamaño fijo.
- Un slice bidimensional (`[][]Tipo`) es dinámico en cuanto a las dimensiones y puede cambiar de tamaño durante la ejecución.
- Las matrices pueden tener diferentes números de columnas en cada fila, lo que se llama
"matrices irregulares" o "jagged arrays".
- Para trabajar con las matrices en Go, debemos usar bucles anidados para recorrerlas.

* Ejemplos:
- Una matriz de tamaño fijo de 3x3 con arreglos se define así: `va matriz [3][3]int`.
- Una matriz dinámica o slice bidimensional de 3x3 se define como: `matriz := make([][]int, 3)` y luego se inicializan las columnas.

* Indices:
- Las matrices están `indexadas` desde 0, al igual que los slices. Es decir, la primera fila y columna
comienzan en el índice 0.
- El acceso a los elementos de una matriz se hace a través de dos índices: [fila][columna].
- Se pueden modificar los elementos directamente asignando nuevos valores usando los índices.

* Filas y Columnas:
- Filas: Son secuencias horizontales de elementos.
- Columnas: Son secuencias verticales de elementos.

* Ejemplo visual:
[
  [1, 2, 3],  // Fila 0
  [4, 5, 6],  // Fila 1
  [7, 8, 9]   // Fila 2
]
En este caso, el valor en la posición [1][2] es 6 (fila 1, columna 2).
*/

func main() {
	// Ejemplo: Matriz de 1x1 (1 fila, 1 columna).
	matriz := [][]int{
		{1},
	}
	fmt.Println("Matriz 1x1:", matriz)                        // Output: [[1]]
	fmt.Println("Elemento en posición [0][0]:", matriz[0][0]) // Output: 1

	// Ejemplo: Matriz de 2x2 (2 filas, 2 columnas).
	matriz2 := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("Matriz 2x2:", matriz2)                        // Output: [[1, 2], [3, 4]]
	fmt.Println("Elemento en posición [1][1]:", matriz2[1][1]) // Output: 4

	// Ejemplo: Matriz de 3x3 (3 filas, 3 columnas).
	matriz3 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matriz 3x3:", matriz3)                        // Output: [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
	fmt.Println("Elemento en posición [2][0]:", matriz3[2][0]) // Output: 7

	// Ejemplo: Matriz de 4x4 (4 filas, 4 columnas).
	matriz4 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println("Matriz 4x4:", matriz4)
	fmt.Println("Elemento en posición [3][3]:", matriz4[3][3]) // Output: 16

	// Ejemplo: Matriz irregular (jagged array).
	// 4 filas, la primera fila tiene 1 columna, y las demás tienen 3 columnas.
	matriz5 := [][]int{
		{3},          // Fila 0: 1 elemento
		{11, 2, 4},   // Fila 1: 3 elementos
		{4, 5, 6},    // Fila 2: 3 elementos
		{10, 8, -12}, // Fila 3: 3 elementos
	}
	fmt.Println("Matriz irregular:", matriz5)                  // Output: [[3] [11 2 4] [4 5 6] [10 8 -12]]
	fmt.Println("Elemento en posición [3][2]:", matriz5[3][2]) // Output: -12

	//* Recorrido de una matriz usando bucles anidados.
	fmt.Println("Recorriendo matriz3:")
	for i := range matriz3 { // Iteramos sobre las filas
		for j := 0; j < len(matriz3[i]); j++ { // Iteramos sobre las columnas de cada fila
			fmt.Printf("Elemento [%d][%d] = %d\n", i, j, matriz3[i][j])
		}
	}

	//* Recorremos la matriz usando dos bucles anidados `for range`.
	// Usamos como ejemplo la matriz 2x2 de arriba.
	for i, fila := range matriz2 {
		for j, valor := range fila {
			fmt.Printf("Elemento en posición [%d][%d] = %d\n", i, j, valor)
		}
	}

	// Crear una matriz de tamaño dinámico (n x m).
	n, m := 3, 4                       // Definimos una matriz de 3 filas y 4 columnas
	matrizDinamica := make([][]int, n) // Creamos el slice de las filas
	for i := range matrizDinamica {
		matrizDinamica[i] = make([]int, m) // Asignamos a cada fila un slice de columnas
	}

	// Llenar la matriz con valores (por ejemplo, el producto de sus índices).
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrizDinamica[i][j] = i * j
		}
	}
	fmt.Println("Matriz dinámica 3x4:", matrizDinamica) // Output: [[0 0 0 0] [0 1 2 3] [0 2 4 6]]

	// Matriz de tipo float64 (soporte para otros tipos de datos).
	matrizFloat := [][]float64{
		{1.1, 2.2, 3.3},
		{4.4, 5.5, 6.6},
	}
	fmt.Println("Matriz de floats:", matrizFloat)
}
