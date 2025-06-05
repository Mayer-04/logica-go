package main

import "fmt"

/*
* Bucle For: Para
- El bucle `for` se utiliza para repetir un bloque de código mientras una condición sea verdadera.
- En Go solo existe un tipo de bucle `for``, pero puede usarse de diferentes maneras.
- Sintaxis básica: for `inicialización`; `condición`; `actualización` { bloque de código }
- Se puede utilizar para `iterar` sobre colecciones como arrays, slices, strings, maps y canales.

* Break: Romper
- Termina el bucle completamente y sigue con el código después del bloque `for`.
* Continue: Continuar
- Salta el resto de la iteración actual y pasa a la siguiente.
- Saltar: Omitir lo que queda en esa 'vuelta' del bucle y pasar directamente a la siguiente.
- No sale del bucle, solo omite el código restante de la iteración actual.

* Conceptos claves:
- Recorrer o Iterar: Es pasar por cada uno de los elementos de una colección de datos, como en un ciclo.
- Iterable: Es cualquier estructura de datos que se puede recorrer o iterar.
- Iteración: Es el proceso de repetir una acción varias veces. Cada vez que se pasa sobre un elemento de la colección
es una iteración, y el proceso continúa hasta que no queden más elementos.
*/

func main() {
	//* Bucle "for" clásico.
	// Sintaxis: inicialización; condición; actualización.
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//* Bucle tipo "while".
	// Go no tiene un bucle "while", pero puede simularse con "for".
	i := 0
	for i <= 3 {
		fmt.Println("Bucle tipo while:", i)
		i++
	}

	//* Bucle "for" con "range":
	// Se usa para recorrer colecciones como arrays, slices, strings, mapas o canales.
	// Puedes obtener el índice y el valor en cada iteración.
	// Si solo necesitas el valor, omite el índice usando `_` (por ejemplo: _, v).
	// Si solo necesitas el índice, omite el valor directamente.
	numeros := [4]int{1, 2, 3, 4}
	// En este ejemplo, "i" es el índice y "v" es el valor de cada elemento.
	for i, v := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", i, v)
	}

	//* Bucle "for" con rango de enteros (Go 1.22+):
	// Permite usar "range" con enteros para iterar desde 0 hasta n-1.
	for i := range 3 {
		fmt.Println(i) // Salida: 0, 1, 2
	}

	//* Bucle for con rango de enteros sin variable de índice.
	// Ignoramos la variable índice `i`.
	// Es útil cuando necesitas realizar una acción repetitiva un número determinado de veces.
	for range 5 {
		fmt.Println("Go") // Output: Go Go Go Go Go
	}

	//* Bucle for con "break".
	for i := 1; i <= 50; i++ {
		// Cuando i sea igual a 25 se ejecuta el `break`, rompe el bucle.
		if i == 25 {
			break
		}
		// Si la condición del if no es verdadera, el código sigue y ejecuta `fmt.Println(i)`.
		// fmt.Println(i) solo se ejecuta si el valor de `i` es menor a 25.
		fmt.Println(i) // Imprime los números del 1 al 24.
	}

	//* Bucle for con "continue".
	for i := 1; i <= 5; i++ {
		// Si i es igual a 4, imprime "Continuando..." y pasa directamente a la siguiente iteración,
		// sin ejecutar el código que imprime el número.
		if i == 4 {
			fmt.Println("Continuando...")
			continue
		}
		// Para todos los valores de i que no sean 4, imprime "Soy el número" seguido del valor de i.
		fmt.Printf("Soy el número %d\n", i)
	}

	//* Iterar sobre un string.
	// Se itera sobre runas (runes), que son valores Unicode.
	// Debes convertir cada rune a string para imprimirlo como un carácter.
	nombre := "Mayer"
	for _, value := range nombre {
		fmt.Println(string(value))
	}

	//* Bucle infinito.
	// Un bucle sin condición explícita se convierte en un bucle infinito.
	count := 0
	for {
		fmt.Println("Bucle infinito, cuenta:", count)
		// Cuando count sea igual a 3, se sale del bucle.
		if count == 3 {
			break
		}
		count++
	}

	// Matrices
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	//* Bucles anidados para trabajar con matrices (slice de slices).
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("Elemento en [%d][%d] es %d\n", i, j, matrix[i][j])
		}
	}

	//* Bucle "for" con múltiples variables.
	// Se inicializan dos variables: i = 0 y j = 10.
	// El bucle se ejecuta mientras i < j (i sea menor que j).
	// En cada iteración, i se incrementa y j se decrementa.
	// Imprime i de 0 a 4 y j de 10 a 6.
	for i, j := 0, 10; i < j; i, j = i+1, j-1 {
		fmt.Printf("i: %d, j: %d\n", i, j)
	}

	//* Simulación de un bucle "do while".
	// El código se ejecuta al menos una vez antes de la verificación de la condición.
	var numero int
	for {
		// Código que se ejecuta al menos una vez
		fmt.Print("Introduce un número positivo: ")
		// fmt.Scan -> Permite leer la entrada del usuario
		fmt.Scan(&numero)

		// Condición para romper el bucle en caso de que el número sea negativo con `break`
		if numero > 0 {
			break
		}
		fmt.Println("El número no es positivo. Inténtalo de nuevo.")
	}

	fmt.Printf("Número válido introducido: %d\n", numero)
}
