package main

import (
	"fmt"
	"slices"
)

/*
* Slices: Rebanadas o Cortes
- Los slices son estructuras más flexibles que los arrays, con longitud dinámica.
- A diferencia de los arrays, los slices pueden cambiar de tamaño durante la ejecución del programa.
- La estructura de un slice (puntero, longitud, capacidad) generalmente se asigna en la pila si es seguro hacerlo,
  pero el array subyacente puede ser asignado en el montón si es necesario.
- La longitud de un slice es el número de elementos que contiene.
* Capacidad de un slice:
- Es el número de elementos desde el inicio del slice hasta el final del array subyacente.
- Cuando se excede la capacidad del slice, Go crea un nuevo array, generalmente duplicando la capacidad,
  pero este comportamiento puede variar dependiendo del tamaño.

Un slice internamente es una estructura con tres campos:

type slice struct {
	array unsafe.Pointer // Puntero al array subyacente
	len   int            // Longitud del slice
	cap   int            // Capacidad del array
}
*/

func main() {
	// Declaración de un slice.
	// Cuando se declara un slice sin inicializarlo (sin asignarle valores o longitud), su valor por defecto es nil.
	var slice []int
	fmt.Println("slice sin valores:", slice)

	// Creando un slice nulo explícitamente.
	// Un slice nulo es igual a nil y no contiene elementos.
	sliceNull := []string(nil)
	fmt.Println("slice nulo:", sliceNull)

	// Declaración literal de un slice.
	// Aquí se declaran e inicializan los valores del slice en una sola línea.
	slice2 := []int{1, 2, 3, 4}
	fmt.Println("slice literal:", slice2)

	// Declaración literal de un slice vacío.
	// Los slices vacíos tienen una longitud de 0 y su capacidad es 0.
	// Debe evitarse si inicializamos un slice sin elementos.
	sliceEmpty := []int{}
	fmt.Println("slice literal vacío:", sliceEmpty)

	// Declaración de un slice utilizando la función new.
	// new([]int) devuelve un puntero a un slice vacío. Desreferenciándolo obtenemos un slice vacío pero no nil.
	sliceNew := *new([]int)
	fmt.Println("slice new:", sliceNew) // Output: []

	// Declaración literal de un slice con índices explícitos.
	// El slice resultante no mantiene el orden de declaración.
	// Los índices pueden ser especificados explícitamente para definir elementos en posiciones específicas.
	chars := []string{0: "a", 2: "m", 1: "z"}
	fmt.Println("slice con índices explícitos:", chars) // Output: ["a", "z", "m"]

	// Crear un slice a partir de un array.
	// Se usa el slicing (:) para crear un nuevo slice a partir de un array existente.
	// Sintaxis: array[start:end]
	array := [5]int{1, 2, 3, 4, 5}
	slice3 := array[1:3]
	fmt.Println("slice 3 (array[1:3]):", slice3) // Output: [2 3]

	// Crear un slice desde una posición hasta el final del array.
	// El slice incluye todos los elementos desde el índice especificado hasta el final del array.
	slice4 := array[1:]
	fmt.Println("slice 4 (array[1:]):", slice4) // Output: [2 3 4 5]

	//* Funciones integradas para trabajar con slices

	// Tamaño o longitud de un slice.
	// La longitud del slice se obtiene con la función `len()`.
	fmt.Println("longitud de slice4:", len(slice4))

	// Capacidad de un slice.
	// La capacidad del slice, que es el número de elementos desde el inicio del slice hasta el final del array
	// subyacente, se obtiene con `cap()`.
	fmt.Println("capacidad de slice4:", cap(slice4))

	// Añadir elementos a un slice.
	// La función `append()` agrega uno o más elementos al slice y modifica el `array subyacente`.
	slice5 := append(slice4, 6)
	fmt.Println("slice después de append:", slice5)

	// Limpiar un slice con la función `clear()`.
	// `clear()` pone los elementos a su valor cero pero mantiene la longitud y capacidad del slice.
	clear(slice4)
	fmt.Println("slice después de clear:", slice4)

	// Desempaquetar un slice (...).
	// Los tres puntos permiten expandir un slice y pasarlo como argumentos individuales a una función.
	a := []int{1, 2, 3, 4}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println("slice después de desempaquetar:", c)

	// Crear un slice la función `make()`.
	// Permite crear un slice con una longitud y capacidad inicial especificadas.
	// La capacidad es un argumento opcional.
	// El slice creado se inicializa con valores cero de su tipo correspondiente.
	makeBytes := make([]byte, 5)
	fmt.Println("slice creado con make:", makeBytes)

	// Función `Clone`.
	// Crea una copia superficial de un slice.
	sliceToClone := []int{0, 42, -10, 8}
	clonedSlice := slices.Clone(sliceToClone)
	fmt.Println("clone de slice:", clonedSlice)

	// Función `Contains`.
	// Verifica si un slice contiene un elemento. Devuelve `true` o `false`.
	sliceToSearch := []int{0, 1, 2, 3}
	elementPresent := slices.Contains(sliceToSearch, 2)
	fmt.Println("contains:", elementPresent) // Output: true

	// Función `Delete`.
	// Elimina un rango de elementos en un slice, especificando el índice de inicio y fin (exclusivo).
	sliceToDelete := []string{"a", "b", "c", "d", "e"}
	deletedSlice := slices.Delete(sliceToDelete, 1, 3)
	fmt.Println("slice después de delete:", deletedSlice) // Output: [a e]

	// Función `Equal`.
	// Verifica si dos slices son iguales.
	// Devuelve `true` si ambos slices tienen los mismos elementos en el mismo orden.
	sliceEqual := []int{0, 42, 8}
	fmt.Println(slices.Equal(sliceEqual, []int{0, 42, 8})) // Output: true
	fmt.Println(slices.Equal(sliceEqual, []int{10}))       // Output: false

	// Función `Index`.
	// Busca un elemento en un slice y devuelve su índice. Devuelve `-1` si el elemento no se encuentra.
	sliceToIndex := []int{0, 42, 8}
	fmt.Println("índice de 42:", slices.Index(sliceToIndex, 42)) // Output: 1

	// Función `Clip`.
	// Se utiliza para eliminar la capacidad no utilizada de un slice.
	// Liberar esta capacidad extra no utilizada para reducir el uso de memoria.
	fmt.Println("capacidad del slice c:", cap(c)) // Output: 8
	sliceClip := slices.Clip(c)
	fmt.Println("capacidad del nuevo slice con clip:", cap(sliceClip)) // Output: 7

	// Función `Grow`.
	// Aumenta la capacidad de un slice sin modificar su longitud actual.
	// Se puede especificar el número de elementos adicionales que se desea agregar.
	sliceToGrow := []int{0, 42, 8}
	grownSlice := slices.Grow(sliceToGrow, 2)
	fmt.Println("slice después de grow:", grownSlice)
	fmt.Println("longitud:", len(grownSlice))
	fmt.Println("capacidad:", cap(grownSlice))

	// Función `Sort`.
	// Ordena los elementos en un slice en orden ascendente.
	sliceToSort := []int{0, 42, 8, -10, 2}
	slices.Sort(sliceToSort)
	fmt.Println("slice ordenado:", sliceToSort) // Output: [-10 0 2 8 42]

	// Función `Concat` (Introducida en Go 1.22).
	// Concatena dos o más slices en un nuevo slice.
	s1 := []string{"Mayer", "Andres"}
	s2 := []string{"Joe", "Rose"}
	concat := slices.Concat(s1, s2)
	fmt.Println("slice concatenado:", concat) // Output: [Mayer Andres Joe Rose]

	// Función `Repeat` (Introducida en Go 1.23).
	// Repite un slice un número determinado de veces para crear un nuevo slice más grande.
	numbers := []int{0, 1, 2, 3}
	repeat := slices.Repeat(numbers, 2)
	fmt.Println("slice repetido:", repeat) // Output: [0 1 2 3 0 1 2 3]
}
