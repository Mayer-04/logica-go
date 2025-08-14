package main

import (
	"fmt"
	"slices"
)

/*
* Slices en Go: Rebanadas o Cortes
Un `slice` es una estructura que permite trabajar con colecciones de elementos cuya longitud puede cambiar.
A diferencia de los arrays, los slices son más flexibles y se usan ampliamente en Go.

- Se declaran con el tipo de dato `[]T`, donde `T` es el tipo de dato de los elementos del slice.
- Longitud: Es el número de elementos que contiene.

Internamente, un slice contiene:
- Un puntero a un array subyacente
- Su longitud actual
- Su capacidad (cuántos elementos puede contener antes de necesitar crecer)

Estructura simplificada:

type slice struct {
	array unsafe.Pointer // Puntero al array subyacente
	len   int            // Longitud del slice
	cap   int            // Capacidad del array
}

📌 Cuando se supera la capacidad del slice, Go crea un nuevo array con mayor capacidad (usualmente el doble).
*/

func main() {
	// Declaración de un slice sin inicialización.
	// Un slice declarado pero no inicializado tiene el valor por defecto nil.
	// No apunta a ningún array subyacente y no contiene elementos.
	var slice []int
	fmt.Println("slice sin valores:", slice)

	// Creación explícita de un slice nulo.
	// Asignar nil directamente deja claro que el slice no apunta a ninguna estructura subyacente.
	sliceNull := []string(nil)
	fmt.Println("slice nulo:", sliceNull)

	// Declaración literal de un slice.
	// Aquí se declaran e inicializan los valores del slice en una sola línea.
	slice2 := []int{1, 2, 3, 4}
	fmt.Println("slice literal:", slice2)

	// Declaración literal de un slice vacío.
	// Tiene longitud y capacidad igual a 0, pero a diferencia de un slice nil, sí está inicializado.
	// Debe evitarse si inicializamos un slice sin elementos.
	sliceEmpty := []int{}
	fmt.Println("slice literal vacío:", sliceEmpty)

	// Declaración de un slice utilizando la función new.
	// new([]int) devuelve un puntero a un slice vacío. Desreferenciándolo obtenemos un slice vacío pero no nil.
	sliceNew := *new([]int)
	fmt.Println("slice new:", sliceNew) // Output: []

	// Slice con índices explícitos.
	// El slice resultante no mantiene el orden de declaración.
	// Los índices pueden ser especificados explícitamente para definir elementos en posiciones específicas.
	chars := []string{0: "a", 2: "m", 1: "z"}
	fmt.Println("slice con índices explícitos:", chars) // Output: ["a", "z", "m"]

	// Creación de un slice a partir de un array usando slicing.
	// Sintaxis: array[inicio:fin] (incluye el índice inicio, excluye el fin).
	array := [5]int{1, 2, 3, 4, 5}
	slice3 := array[1:3]
	fmt.Println("slice 3 (array[1:3]):", slice3) // Output: [2 3]

	// Slice desde una posición hasta el final del array.
	// Cuando se omite el índice final, se asume hasta el último elemento.
	slice4 := array[1:]
	fmt.Println("slice 4 (array[1:]):", slice4) // Output: [2 3 4 5]

	//* Funciones integradas para trabajar con slices

	// Tamaño o longitud de un slice.
	// La longitud del slice se obtiene con la función `len()`.
	fmt.Println("longitud de slice4:", len(slice4)) // Output: 4

	// Capacidad de un slice.
	// La capacidad del slice, que es el número de elementos desde el inicio del slice hasta el final del array
	// subyacente, se obtiene con `cap()`.
	fmt.Println("capacidad de slice4:", cap(slice4)) // Output: 4

	// Añadir elementos a un slice.
	// La función `append()` agrega uno o más elementos al slice y modifica el `array subyacente`.
	slice5 := append(slice4, 6)
	fmt.Println("slice después de append:", slice5) // Output: [2 3 4 5 6]

	// Agregar múltiples elementos a un slice.
	slice6 := append(slice4, 6, 7, 8)
	fmt.Println("agregando múltiples elementos:", slice6) // Output: [2 3 4 5 6 7 8]

	// Limpiar un slice con la función `clear()`.
	// `clear()` pone los elementos a su valor cero pero mantiene la longitud y capacidad del slice.
	clear(slice4)
	fmt.Println("slice después de clear:", slice4) // Output: [0 0 0 0]

	// Uso del operador de desempaquetado (...) con slices en Go.
	// El operador "..." permite expandir un slice para que sus elementos se pasen como argumentos individuales.
	// Esto es útil, por ejemplo, al usar la función append para combinar múltiples slices en uno solo.
	a := []int{1, 2, 3, 4}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println("Slice después de usar el operador de desempaquetado:", c) // Output: [1 2 3 4 4 5 6]

	// Crear un slice usando la función `make()`.
	// `make` permite crear un slice con una longitud inicial y opcionalmente una capacidad.
	// El slice creado se inicializa con valores cero de su tipo correspondiente.
	// En este caso, se crea un slice de 5 bytes, todos con valor 0.
	makeBytes := make([]byte, 5)
	fmt.Println("slice creado con make:", makeBytes) // Output: [0 0 0 0 0]

	// Crear un slice usando `make()` especificando longitud y capacidad.
	// La longitud indica cuántos elementos están inicialmente presentes en el slice.
	// La capacidad define cuántos elementos puede contener antes de que se necesite reasignar memoria.
	// En este ejemplo, el slice tiene una longitud de 3 y una capacidad de 5.
	// Los elementos iniciales se llenan con ceros (valor por defecto del tipo `int`).
	numbers := make([]int, 3, 5)
	fmt.Println("slice:", numbers)          // Output: [0 0 0]
	fmt.Println("longitud:", len(numbers))  // Output: 3
	fmt.Println("capacidad:", cap(numbers)) // Output: 5

	// Función `Clone`.
	// Crea una copia superficial de un slice.
	sliceToClone := []int{0, 42, -10, 8}
	clonedSlice := slices.Clone(sliceToClone)
	fmt.Println("clone de slice:", clonedSlice) // Output: [0 42 -10 8]

	// Función `Contains`.
	// Verifica si un slice contiene un elemento. Devuelve `true` o `false`.
	sliceToSearch := []int{0, 1, 2, 3}
	elementPresent := slices.Contains(sliceToSearch, 2)
	fmt.Println("contains:", elementPresent) // Output: true

	// Función `Compare` (Go 1.21).
	// Compara dos slices elemento por elemento:
	// Devuelve `0` si los slices son iguales.
	// Devuelve `1` si el primer slice es mayor que el segundo.
	// Devuelve `-1` si el primer slice es menor que el segundo.
	names := []string{"Alice", "Bob", "Vera"}
	referenceNames := []string{"Alice", "Bob", "Vera"}
	fmt.Println("Equal:", slices.Compare(names, referenceNames)) // Output: 0

	// EJEMPLO 2
	firstNumbers := []int{1, 2, 3}
	secondNumbers := []int{1, 2, 4}
	fmt.Println(slices.Compare(firstNumbers, secondNumbers)) // Output: -1 (porque 3 < 4)

	// Función `Delete`.
	// Elimina un rango de elementos en un slice, especificando el índice de inicio y fin (exclusivo).
	sliceToDelete := []string{"a", "b", "c", "d", "e"}
	deletedSlice := slices.Delete(sliceToDelete, 1, 3)
	fmt.Println("slice después de delete:", deletedSlice) // Output: [a d e]

	// Función `DeleteFunc`.
	// Elimina los elementos de un slice según una condición personalizada que defines con una función.
	// Es como decir: “Quiero borrar todos los elementos que cumplan cierta regla”.
	fibonacciNumbers := []int{0, 1, 1, 2, 3, 5, 8}
	fibonacciNumbers = slices.DeleteFunc(fibonacciNumbers, func(number int) bool {
		return number%2 != 0 // Elimina los impares
	})

	fmt.Println(fibonacciNumbers) // Output: [0 2 8]

	// Función `Equal`.
	// Verifica si dos slices son iguales en longitud y elementos.
	// Devuelve `true` si ambos slices tienen los mismos elementos en el mismo orden.
	sliceEqual := []int{0, 42, 8}
	fmt.Println(slices.Equal(sliceEqual, []int{0, 42, 8})) // Output: true
	fmt.Println(slices.Equal(sliceEqual, []int{10}))       // Output: false

	// Función `Index`.
	// Busca un elemento en un slice y devuelve su índice.
	// Devuelve `-1` si el elemento no se encuentra.
	sliceToIndex := []int{0, 42, 8}
	fmt.Println("índice de 42:", slices.Index(sliceToIndex, 42)) // Output: 1

	// Función `Clip`.
	// Se utiliza para eliminar la capacidad no utilizada de un slice.
	// Liberar esta capacidad extra no utilizada reduce el uso de memoria.
	fmt.Println("capacidad del slice c:", cap(c)) // Output: 8
	sliceClip := slices.Clip(c)
	fmt.Println("capacidad del nuevo slice con clip:", cap(sliceClip)) // Output: 7

	// Función `Grow`.
	// Aumenta la capacidad de un slice sin modificar su longitud actual.
	// Se puede especificar el número de elementos adicionales que se desea agregar.
	sliceToGrow := []int{0, 42, 8}
	grownSlice := slices.Grow(sliceToGrow, 2)
	fmt.Println("slice después de grow:", grownSlice) // Output: [0 42 8]
	fmt.Println("longitud:", len(grownSlice))         // Output: 3
	fmt.Println("capacidad:", cap(grownSlice))        // Output: 6

	// Función `Sort`.
	// Ordena los elementos en un slice en orden ascendente.
	// Modifica el slice original.
	sliceToSort := []int{0, 42, 8, -10, 2}
	slices.Sort(sliceToSort)
	fmt.Println("slice ordenado:", sliceToSort) // Output: [-10 0 2 8 42]

	// Función `Compact`.
	// Elimina los elementos duplicados `consecutivos` en un slice.
	// No elimina todos los duplicados, solo los elementos que están seguidos (uno tras otro).
	// Modifica el slice original.
	words := []string{"go", "go", "lang", "lang", "lang", "rocks"}
	words = slices.Compact(words)
	fmt.Println(words)

	// Función `Concat` (Introducida en Go 1.22).
	// Concatena dos o más slices en un nuevo slice.
	s1 := []string{"Mayer", "Andres"}
	s2 := []string{"Joe", "Rose"}
	concat := slices.Concat(s1, s2)
	fmt.Println("slice concatenado:", concat) // Output: [Mayer Andres Joe Rose]

	// Función `Repeat` (Introducida en Go 1.23).
	// Repite un slice un número determinado de veces para crear un nuevo slice más grande.
	numbers2 := []int{0, 1, 2, 3}
	repeat := slices.Repeat(numbers2, 2)
	fmt.Println("slice repetido:", repeat) // Output: [0 1 2 3 0 1 2 3]
}
