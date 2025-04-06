package main

import "fmt"

/*
* Arrays en Go: Arreglos
Los Arrays en Go son estructuras de datos que tienen un tamaño fijo y almacenan elementos del mismo tipo
en ubicaciones de memoria contiguas.

- Los arrays suelen almacenarse en la pila (stack).
- El tamaño del array es conocido en tiempo de compilación y es parte de su tipo.
	Ejemplo: `[5]int` y `[3]int` son tipos distintos y no compatibles entre sí.
- `len(arr)` y `cap(arr)` son constantes y siempre iguales al tamaño del array.
- Los elementos del array se acceden mediante índices,
donde el primer elemento está en `arr[0]` y el último en `arr[len(arr)-1]`.
- La dirección de memoria del array es la misma que la del primer elemento (`arr[0]`),
y los elementos están almacenados de forma contigua en memoria.
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
	var numbers [5]int
	// Asignación de valores a los elementos del array.
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("array de números:", numbers) // Output: [10 20 30 40 50]

	//* IMPORTANTE: Los arrays son secuencias contiguas de elementos almacenados uno tras otro en la memoria.
	byteArray := [5]byte{0, 1, 2, 3, 4}
	println("byteArray", &byteArray)       // Output: 0xc000027b71
	println("byteArray[0]", &byteArray[0]) // Output: 0xc000027b71
	println("byteArray[1]", &byteArray[1]) // Output: 0xc000027b72
	println("byteArray[2]", &byteArray[2]) // Output: 0xc000027b73
	println("byteArray[3]", &byteArray[3]) // Output: 0xc000027b74
	println("byteArray[4]", &byteArray[4]) // Output: 0xc000027b75

	// Ejemplo de uso de la función `len()` para obtener la longitud del array.
	fmt.Println("longitud del array de números:", len(numbers)) // Output: 5

	// Declaración de un array con valores literales.
	// Los valores se proporcionan directamente entre llaves.
	namesArray := [3]string{"Mayer", "Andres", "Chaves"}
	fmt.Println("array literal de nombres:", namesArray) // Output: [Mayer Andres Chaves]

	// * IMPORTANTE:
	// Cuando declaramos un array literal: arr := [3]int{1, 2, 3, 4}, lo que realmente sucede es esto:
	arr := [4]int{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4

	// Inicialización basada en índices.
	indexedArray := [3]int{0: 1, 2: 3}
	fmt.Println("array con inicialización basada en índices:", indexedArray) // Output: [1 0 3]
	a := [...]int{0: 1, 2: 3, 1}
	fmt.Println("array con longitud implícita basada en índices:", a) // Output: [1 0 3 1]

	// Array con longitud implícita.
	// El tamaño del array se infiere automáticamente basado en la cantidad de elementos proporcionados.
	array4 := [...]int{10, 20, 30}
	fmt.Printf("array con longitud implícita: %v, longitud: %d\n", array4, len(array4)) // [10 20 30], longitud: 3

	// * Array con longitud implícita - Inicialización indexada.
	// Se inicializan elementos específicos de un array usando índices,
	// mientras que los elementos no especificados se inicializan con el valor por defecto.
	// En este caso, el tamaño del array será deducido automáticamente.
	// El array tendrá una longitud de 6 elementos (de 0 a 5), el índice 5 se inicializa con el valor 3.
	indexedImplicitArray := [...]int{5: 3}
	fmt.Println("array con índice 5 asignado:", indexedImplicitArray) // Output: [0 0 0 0 0 3]

	// Acceder a un elemento del array.
	// Los índices comienzan en 0, por lo que namesArray[1] accede al segundo elemento.
	fmt.Println("segundo elemento de namesArray:", namesArray[1]) // Output: Andres

	// Modificación de un elemento del array.
	// El primer elemento del array se cambia de "Perro" a "Conejo".
	animals := [3]string{"Perro", "Gato", "Pez"}
	animals[0] = "Conejo"
	fmt.Println("array modificado:", animals) // Output: [Conejo Gato Pez]

	// Los arrays en Go se pasan por valor.
	// Asignar un array a otra variable crea una `copia` de todos los elementos del array original.
	copyArray := array
	fmt.Println("copia del array:", copyArray) // Output: [false false]

	// Declaración de un array multidimensional (una matriz).
	// Se definen arrays dentro de arrays para crear una estructura de datos bidimensional.
	matrix := [3][2]int{{10, 20}, {30, 40}, {50, 60}}
	fmt.Println("matriz:", matrix) // Output: [[10 20] [30 40] [50 60]]

	// Los arrays se pueden comparar con `==` si tienen el mismo tipo, longitud y los elementos son iguales.
	// Como `array` y `copyArray` tienen la misma longitud y los mismos elementos, son iguales.
	fmt.Println("comparando arrays:", array == copyArray) // Output: true

	// Recorriendo un array con un for `range`.
	names := [3]string{"Mayer", "Andres", "Luis"}

	for i := range len(names) {
		fmt.Println("elemento:", names[i])
	}

	// Recorriendo un array con un for `range` con su índice y su valor.
	for index, value := range names {
		fmt.Println("índice:", index, "valor:", value)
	}
}
