package main

import "fmt"

/*
* Arrays: Arreglos
Los Arrays en Go son estructuras de datos que tienen un tamaño fijo y almacenan elementos del mismo tipo
en ubicaciones de memoria contiguas.

- Los arrays suelen declarar en la pila (stack).
- El compilador de Go conoce la longitud y la capacidad con anticipación de un array.
- Para un array, len(arr) y cap(arr) son constantes y siempre iguales al tamaño del array.
- Los elementos del array se acceden mediante `índices`, donde el primer elemento está en el índice 0
y el último está en `len(arr)-1`.
- La dirección en memoria de un array es la misma que la dirección del primer elemento (arr[0]).
- Cuando pasas un array a una función, se crea una copia del array. Esto evita modificaciones no deseadas
en el array original dentro de la función.
- No se puede usar "slicing" directamente en arrays. El slicing es una característica de los slices y strings.
- Go soporta arrays multidimensionales, que son útiles para representar matrices o tablas.
- Se pueden comparar dos arrays si tienen el mismo tipo y longitud.
*/

func main() {
	// Array sin asignación de valores.
	// Los elementos del array se inicializan con el valor cero del tipo de dato correspondiente.
	// En este caso array tiene 2 elementos que corresponden al valor cero del tipo de dato `bool`.
	var array [2]bool
	fmt.Println("array sin asignación de valores:", array) // Output: [false false]

	// Declaración de un array de enteros con una longitud de 5.
	var array1 [5]int
	// Asignación de valores a los elementos del array.
	array1[0] = 10
	array1[1] = 20
	array1[2] = 30
	array1[3] = 40
	array1[4] = 50
	fmt.Println("array:", array) // Output: [10 20 30 40 50]

	// Ejemplo de uso de la función `len()` para obtener la longitud del array.
	fmt.Println("longitud de array:", len(array)) // Output: 5

	// Declaración de un array con valores literales.
	// Los valores se proporcionan directamente entre llaves.
	array2 := [3]string{"Mayer", "Andres", "Chaves"}
	fmt.Println("array literal:", array2)

	// * Cuando declaramos un array literal: arr := [3]int{1, 2, 3, 4}, lo que realmente sucede es esto:
	arr := [4]int{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	// Declaración de un array multidimensional.
	// Se definen arrays dentro de arrays para crear una estructura de datos bidimensional.
	array3 := [3][2]int{{10, 20}, {30, 40}, {50, 60}}
	fmt.Println("array multidimensional:", array3)

	// Array con longitud implícita.
	// El tamaño del array se infiere automáticamente basado en la cantidad de elementos proporcionados.
	array4 := [...]int{10, 20, 30}
	fmt.Println("array con longitud implícita:", array4)

	// * Array con longitud implícita - Inicialización indexada.
	// Se inicializan elementos específicos de un array usando índices,
	// mientras que los elementos no especificados se inicializan con el valor por defecto.
	// En este caso, el tamaño del array será deducido automáticamente.
	// El array tendrá una longitud de 6 elementos (de 0 a 5), el índice 5 se inicializa con el valor 3.
	array6 := [...]int{5: 3}
	fmt.Println("array6:", array6) // Output: [0 0 0 0 0 3]

	// Acceder a un elemento del array.
	// Los índices comienzan en 0, por lo que array2[1] accede al segundo elemento.
	fmt.Println("segundo elemento de array2:", array2[1]) // Output: Andres

	// Modificación de un elemento del array.
	// El primer elemento del array se cambia de "Perro" a "Conejo".
	modificar := [3]string{"Perro", "Gato", "Pez"}
	modificar[0] = "Conejo"
	fmt.Println("array modificado:", modificar) // Output: [Conejo Gato Pez]

	// Los arrays en Go se pasan por valor.
	// Asignar un array a otra variable crea una `copia` de todos los elementos del array original.
	var copyArray = array
	fmt.Println("copia del array:", copyArray)

	// Recorriendo un array con un `for clásico`.
	for i := 0; i < len(array); i++ {
		fmt.Println("elemento:", array[i])
	}

	// Recorriendo un array con un for `range`.
	for i, v := range array {
		fmt.Println("i:", i, "v:", v)
	}
}
