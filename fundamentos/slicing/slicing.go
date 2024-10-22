package main

import "fmt"

/*
* Slicing en Go
Es utilizado para obtener una parte de una secuencia de elementos a partir de un array, slice o string.

- `Slicing` es el proceso de crear un nuevo slice a partir de un slice existente.
- El slicing crea un nuevo slice que 'apunta' al mismo array subyacente que el original,
por lo que si modificas uno de los elementos del nuevo slice, estás modificando el array original.
- Puedes aplicar el slicing directamente sobre otro slice.
- El slicing en strings te permite crear una nueva 'subcadena' a partir de una cadena.

* Algunos conceptos útiles:
- Incluido: Significa que ese índice sí forma parte del resultado.
- Excluido: Significa que ese índice no formará parte del resultado. El elemento en esa posición,
no se incluye en el resultado. Por lo que seria algo como el índice - 1.

* Slicing básico: arr[start:end]
- start: índice de inicio (incluido).
- end: índice final (excluido).

* Slicing avanzado: s[start:end:cap]
- Permite especificar un tercer argumento que limita la **capacidad** del nuevo slice.
- `start`: índice de inicio (incluido).
- `end`: índice de fin (excluido).
- `cap`: índice límite de capacidad. La capacidad del nuevo slice será `cap - start`.
NOTA: No es el valor de la capacidad directamente, sino el índice hasta el cual se limita la capacidad.
*/

func main() {
	// Declaración de un array.
	array := [5]int{10, 20, 30, 40, 50}

	// Declaración de un slice a partir de un array.
	// El slice contiene los elementos del array desde el índice 1 hasta el 3 (excluido).
	slice := array[1:3]
	fmt.Println("slice:", slice) // Output: [20 30]

	//* Utilizando valores por defecto.
	// Si omites el valor de `start`, se asume que comienza desde el índice 0.
	slice2 := array[:3]
	// Si omites el valor de `end`, se asume que termina en el último índice del slice original.
	slice3 := array[2:]
	fmt.Println("slice2:", slice2) // Output: [10 20 30]
	fmt.Println("slice3:", slice3) // Output: [30 40 50]

	//* Referencias y no copias.
	// El slicing no crea una copia del slice original, solo apunta al mismo array subyacente.
	original := []string{"a", "b", "c", "d"}
	sliced := original[1:3] // Output: [b, c]
	sliced[0] = "z"
	fmt.Println("original:", original) // Ahora es [a z c d]

	//* Slicing de strings.
	// El slicing te permite crear una nueva 'subcadena', pero no puedes cambiar el contenido de la cadena original.
	// Esto se debe a que los strings son inmutables.
	cadena := "Bienvenidos a Go"
	subcadena := cadena[0:11]
	fmt.Printf("subcadena: %q\n", subcadena) // "Bienvenidos"

	str := "Go es genial"
	substr := str[6:]
	fmt.Printf("substr: %q\n", substr) // "genial"

	//* Slicing avanzado: s[start:end:cap].
	// IMPORTANTE: No es posible aplicar el slicing avanzado a los strings.

	// Declaramos un slice con longitud 5 y capacidad 10.
	s := make([]int, 5, 10)
	fmt.Println("longitud:", len(s), "capacidad:", cap(s)) // Output: longitud: 5, capacidad: 10

	// Usamos s[:len(s):len(s)] para ajustar la capacidad del slice a su longitud.
	// Inicia en el primer elemento, termina en el último elemento y la capacidad es la misma que la longitud.
	s = s[:len(s):len(s)]
	fmt.Println("longitud:", len(s), "capacidad:", cap(s)) // Output: longitud: 5, capacidad: 5

	// Crear un sub-slice desde el índice 1 al 4, pero con capacidad limitada a 3.
	// El tercer parámetro (4) no es el valor de la capacidad, sino el índice límite de la capacidad.
	// Como el sub-slice empieza en el índice 1 y el índice límite es 4, la capacidad es 4 - 1 = 3.
	animales := []string{"perro", "gato", "ballena", "tigre", "loro", "hipopotamo"}
	subSlice := animales[1:4:4]
	fmt.Printf("Sub-slice: %q\n", subSlice)  // Output: ["gato" "ballena" "tigre"]
	fmt.Println("Capacidad:", cap(subSlice)) // Capacidad: 3
}
