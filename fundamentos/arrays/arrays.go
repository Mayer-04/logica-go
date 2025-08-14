package main

import "fmt"

/*
* Arrays en Go: Arreglos

Un array es una colección de elementos del mismo tipo con una longitud fija,
almacenados en ubicaciones contiguas de memoria.

* Características principales:
- El tamaño del array es parte de su tipo, por ejemplo: [5]int.
- Se accede a los elementos por índice, comenzando desde 0.
- Los arrays se almacenan generalmente en la pila (stack).
- La longitud y capacidad del array son constantes y siempre iguales.
- Se pueden comparar entre sí si tienen el mismo tipo y longitud.
- Pasar un array a una función copia todos sus elementos (paso por valor).
- La dirección de memoria del array coincide con la dirección de su primer elemento (arr[0]).
- Los arrays no pueden redimensionarse ni usar "slicing" directamente (eso es para los slices).
- Go permite arrays multidimensionales.
*/

func main() {
	// Declaración de un array sin valores asignados.
	// Los elementos se inicializan con el valor cero de su tipo (false en el caso de bool).
	var array [2]bool
	fmt.Println("array sin asignación de valores:", array) // Output: [false false]

	// Declaración de un array de enteros con longitud 5.
	var numbers [5]int
	// Se asignan valores a cada elemento usando su índice.
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("array de números:", numbers) // Output: [10 20 30 40 50]

	// Los arrays almacenan sus elementos en posiciones contiguas de memoria.
	byteArray := [5]byte{0, 1, 2, 3, 4}
	// Aquí se demuestra accediendo a la dirección de memoria de cada elemento.
	fmt.Println("dirección del array completo:", &byteArray) // Output: 0xc000027b71
	fmt.Println("byteArray[0]", &byteArray[0])               // Output: 0xc000027b71
	fmt.Println("byteArray[1]", &byteArray[1])               // Output: 0xc000027b72
	fmt.Println("byteArray[2]", &byteArray[2])               // Output: 0xc000027b73
	fmt.Println("byteArray[3]", &byteArray[3])               // Output: 0xc000027b74
	fmt.Println("byteArray[4]", &byteArray[4])               // Output: 0xc000027b75

	// Usando len() para obtener la longitud del array.
	fmt.Println("longitud del array de números:", len(numbers)) // Output: 5

	// Declaración de un array con valores literales.
	// Los valores se proporcionan directamente entre llaves ({}) separados por comas.
	namesArray := [3]string{"Mayer", "Andres", "Chaves"}
	fmt.Println("array literal de nombres:", namesArray) // Output: [Mayer Andres Chaves]

	// Cuando declaramos un `array literal arr := [3]int{1, 2, 3}`, lo que realmente sucede es esto:
	arr := [3]int{}
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	// Inicialización basada en índices: se asignan valores a índices específicos.
	indexedArray := [3]int{0: 1, 2: 3}
	fmt.Println("array con inicialización basada en índices:", indexedArray) // Output: [1 0 3]

	// Uso de longitud implícita con inicialización indexada.
	a := [...]int{0: 1, 2: 3, 1}                                      // Se deduce la longitud: 4
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

	// Acceso a elementos por índice.
	fmt.Println("segundo elemento de namesArray:", namesArray[1]) // Output: Andres

	// Modificación de un elemento.
	// El primer elemento del array se cambia de "Perro" a "Conejo".
	animals := [3]string{"Perro", "Gato", "Pez"}
	animals[0] = "Conejo"
	fmt.Println("array modificado:", animals) // Output: [Conejo Gato Pez]

	// Los arrays se pasan por valor, lo que significa que se hace una copia completa.
	copyArray := array
	fmt.Println("copia del array:", copyArray) // Output: [false false]

	// Declaración de un array multidimensional (matriz).
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
